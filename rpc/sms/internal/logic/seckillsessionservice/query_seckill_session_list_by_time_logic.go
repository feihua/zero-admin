package seckillsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySeckillSessionListByTimeLogic 根据时间查询限时购场次
/*
Author: LiuFeiHua
Date: 2025/6/11 10:57
*/
type QuerySeckillSessionListByTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillSessionListByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillSessionListByTimeLogic {
	return &QuerySeckillSessionListByTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillSessionListByTime 根据时间查询限时购场次
func (l *QuerySeckillSessionListByTimeLogic) QuerySeckillSessionListByTime(in *smsclient.QuerySeckillSessionListByTimeReq) (*smsclient.QuerySeckillSessionListByTimeResp, error) {
	sql := `SELECT *
	FROM sms_seckill_session
	WHERE status = 1
	  AND is_deleted = 0
	  AND ? >= start_time
	  AND ? < end_time`

	var result []model.SmsSeckillSession
	db := l.svcCtx.DB
	err := db.WithContext(l.ctx).Raw(sql, in.CurrentTIme, in.CurrentTIme).Scan(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "根据时间查询限时购场次失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据时间查询限时购场次失败")
	}

	var list []*smsclient.SeckillSessionListData
	for _, item := range result {

		list = append(list, &smsclient.SeckillSessionListData{
			Id:         item.ID,                                          // 秒杀场次ID
			Name:       item.Name,                                        // 场次名称
			StartTime:  item.StartTime,                                   // 开始时间
			EndTime:    item.EndTime,                                     // 结束时间
			Status:     item.Status,                                      // 状态：0-禁用，1-启用
			Sort:       item.Sort,                                        // 排序
			CreateBy:   item.CreateBy,                                    // 创建人ID
			CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间
		})
	}

	return &smsclient.QuerySeckillSessionListByTimeResp{
		List: nil,
	}, nil
}
