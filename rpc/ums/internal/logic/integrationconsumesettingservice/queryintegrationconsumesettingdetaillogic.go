package integrationconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationConsumeSettingDetailLogic 查询积分消费设置详情
/*
Author: LiuFeiHua
Date: 2024/6/11 14:23
*/
type QueryIntegrationConsumeSettingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIntegrationConsumeSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingDetailLogic {
	return &QueryIntegrationConsumeSettingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryIntegrationConsumeSettingDetail 查询积分消费设置详情
func (l *QueryIntegrationConsumeSettingDetailLogic) QueryIntegrationConsumeSettingDetail(in *umsclient.QueryIntegrationConsumeSettingDetailReq) (*umsclient.QueryIntegrationConsumeSettingDetailResp, error) {
	q := query.UmsIntegrationConsumeSetting
	item, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询积分消费设置详情失败")
	}

	return &umsclient.QueryIntegrationConsumeSettingDetailResp{
		Id:                 item.ID,                 //
		DeductionPerAmount: item.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: item.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            item.UseUnit,            // 每次使用积分最小单位100
		IsDefault:          item.IsDefault,          // 是否默认：0->否；1->是
		CouponStatus:       item.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
	}, nil

}
