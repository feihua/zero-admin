package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddIntegrationConsumeSettingLogic 添加积分消费设置
/*
Author: LiuFeiHua
Date: 2024/6/11 14:21
*/
type AddIntegrationConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegrationConsumeSettingLogic {
	return &AddIntegrationConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddIntegrationConsumeSetting 添加积分消费设置
func (l *AddIntegrationConsumeSettingLogic) AddIntegrationConsumeSetting(in *umsclient.AddIntegrationConsumeSettingReq) (*umsclient.AddIntegrationConsumeSettingResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {
		q := tx.UmsIntegrationConsumeSetting
		if in.IsDefault == 1 {
			_, err := q.WithContext(l.ctx).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0)
			if err != nil {
				return err
			}
		}
		err := q.WithContext(l.ctx).Create(&model.UmsIntegrationConsumeSetting{
			DeductionPerAmount: in.DeductionPerAmount, // 每一元需要抵扣的积分数量
			MaxPercentPerOrder: in.MaxPercentPerOrder, // 每笔订单最高抵用百分比
			UseUnit:            in.UseUnit,            // 每次使用积分最小单位100
			IsDefault:          in.IsDefault,          // 是否默认：0->否；1->是
			CouponStatus:       in.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
		})

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddIntegrationConsumeSettingResp{}, nil
}
