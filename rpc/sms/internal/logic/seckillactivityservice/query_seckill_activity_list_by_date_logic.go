package seckillactivityservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySeckillActivityListByDateLogic 查询当前时间是否有秒杀活动
/*
Author: LiuFeiHua
Date: 2025/6/11 16:32
*/
type QuerySeckillActivityListByDateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillActivityListByDateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillActivityListByDateLogic {
	return &QuerySeckillActivityListByDateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillActivityListByDate 查询当前时间是否有秒杀活动
func (l *QuerySeckillActivityListByDateLogic) QuerySeckillActivityListByDate(in *smsclient.QuerySeckillActivityListByDateReq) (*smsclient.QueryFlashPromotionListByDateResp, error) {
	currentDate, _ := time.Parse("2006-01-02", in.CurrentDate)
	q := query.SmsSeckillActivity
	result, err := q.WithContext(l.ctx).Where(q.Status.Eq(1), q.StartTime.Lte(currentDate), q.EndTime.Gte(currentDate)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询当前时间是否有秒杀活动失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询当前时间是否有秒杀活动失败")
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

	return &smsclient.QueryFlashPromotionListByDateResp{}, nil
}
