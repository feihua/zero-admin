package member_integration_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationConsumeSettingDetailLogic 查询积分消费设置详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type QueryIntegrationConsumeSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryIntegrationConsumeSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingDetailLogic {
	return &QueryIntegrationConsumeSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryIntegrationConsumeSettingDetail 查询积分消费设置详情
func (l *QueryIntegrationConsumeSettingDetailLogic) QueryIntegrationConsumeSettingDetail(req *types.QueryIntegrationConsumeSettingDetailReq) (resp *types.QueryIntegrationConsumeSettingDetailResp, err error) {

	detail, err := l.svcCtx.IntegrationConsumeSettingService.QueryIntegrationConsumeSettingDetail(l.ctx, &umsclient.QueryIntegrationConsumeSettingDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryIntegrationConsumeSettingDetailData{
		Id:                 detail.Id,                 //
		DeductionPerAmount: detail.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: detail.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            detail.UseUnit,            // 每次使用积分最小单位100
		IsDefault:          detail.IsDefault,          // 是否默认：0->否；1->是
		CouponStatus:       detail.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
	}
	return &types.QueryIntegrationConsumeSettingDetailResp{
		Code:    "000000",
		Message: "查询积分消费设置成功",
		Data:    data,
	}, nil
}
