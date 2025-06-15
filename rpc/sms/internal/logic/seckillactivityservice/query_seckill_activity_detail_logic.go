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
	"gorm.io/gorm"
)

// QuerySeckillActivityDetailLogic 查询秒杀活动详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type QuerySeckillActivityDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillActivityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillActivityDetailLogic {
	return &QuerySeckillActivityDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillActivityDetail 查询秒杀活动详情
func (l *QuerySeckillActivityDetailLogic) QuerySeckillActivityDetail(in *smsclient.QuerySeckillActivityDetailReq) (*smsclient.QuerySeckillActivityDetailResp, error) {
	item, err := query.SmsSeckillActivity.WithContext(l.ctx).Where(query.SmsSeckillActivity.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀活动不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀活动不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀活动异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀活动异常")
	}

	data := &smsclient.QuerySeckillActivityDetailResp{
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
	}

	return data, nil
}
