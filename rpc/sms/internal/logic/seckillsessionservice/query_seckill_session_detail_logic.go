package seckillsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QuerySeckillSessionDetailLogic 查询秒杀场次详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type QuerySeckillSessionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillSessionDetailLogic {
	return &QuerySeckillSessionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillSessionDetail 查询秒杀场次详情
func (l *QuerySeckillSessionDetailLogic) QuerySeckillSessionDetail(in *smsclient.QuerySeckillSessionDetailReq) (*smsclient.QuerySeckillSessionDetailResp, error) {
	item, err := query.SmsSeckillSession.WithContext(l.ctx).Where(query.SmsSeckillSession.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀场次不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀场次不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀场次异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀场次异常")
	}

	data := &smsclient.QuerySeckillSessionDetailResp{
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
	}

	return data, nil
}
