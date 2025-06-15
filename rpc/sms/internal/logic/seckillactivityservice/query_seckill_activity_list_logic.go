package seckillactivityservicelogic

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
	"time"
)

// QuerySeckillActivityListLogic 查询秒杀活动列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type QuerySeckillActivityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillActivityListLogic {
	return &QuerySeckillActivityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillActivityList 查询秒杀活动列表
func (l *QuerySeckillActivityListLogic) QuerySeckillActivityList(in *smsclient.QuerySeckillActivityListReq) (*smsclient.QuerySeckillActivityListResp, error) {
	seckillActivity := query.SmsSeckillActivity
	q := seckillActivity.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(seckillActivity.Name.Like("%" + in.Name + "%"))
	}
	if len(in.StartTime) > 0 {
		startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
		q = q.Where(seckillActivity.StartTime.Gte(startTime))
	}
	if len(in.EndTime) > 0 {
		endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
		q = q.Where(seckillActivity.EndTime.Lte(endTime))
	}
	if in.Status != 2 {
		q = q.Where(seckillActivity.Status.Eq(in.Status))
	}
	if in.IsEnabled != 2 {
		q = q.Where(seckillActivity.IsEnabled.Eq(in.IsEnabled))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀活动列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询秒杀活动列表失败")
	}

	var list []*smsclient.SeckillActivityListData

	for _, item := range result {
		list = append(list, &smsclient.SeckillActivityListData{
			Id:          item.ID,                                          // 编号
			Name:        item.Name,                                        // 活动名称
			Description: item.Description,                                 // 活动描述
			StartTime:   time_util.TimeToStr(item.StartTime),              // 开始时间
			EndTime:     time_util.TimeToStr(item.EndTime),                // 结束时间
			Status:      item.Status,                                      // 状态:0-上线,1-下线
			IsEnabled:   item.IsEnabled,                                   // 是否启用
			CreateBy:    item.CreateBy,                                    // 创建人ID
			CreateTime:  time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:    pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:  time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &smsclient.QuerySeckillActivityListResp{
		Total: count,
		List:  list,
	}, nil
}
