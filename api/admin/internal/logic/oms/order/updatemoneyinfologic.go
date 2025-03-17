package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMoneyInfoLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:05
*/
type UpdateMoneyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMoneyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMoneyInfoLogic {
	return &UpdateMoneyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMoneyInfo 修改订单费用信息
func (l *UpdateMoneyInfoLogic) UpdateMoneyInfo(req *types.UpdateMoneyInfoReq) (resp *types.UpdateMoneyInfoResp, err error) {
	_, err = l.svcCtx.OrderService.UpdateMoneyInfo(l.ctx, &omsclient.UpdateMoneyInfoReq{
		Status:         req.Status,
		OrderId:        req.Id,
		FreightAmount:  req.FreightAmount,
		DiscountAmount: req.DiscountAmount,
		OperateMan:     l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改订单费用信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("修改订单费用信息失败")
	}

	return &types.UpdateMoneyInfoResp{
		Code:    "000000",
		Message: "修改订单费用信息成功",
	}, nil
}
