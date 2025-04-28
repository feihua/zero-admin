package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateReceiverInfoLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:05
*/
type UpdateReceiverInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateReceiverInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReceiverInfoLogic {
	return &UpdateReceiverInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateReceiverInfo 修改收货人信息
func (l *UpdateReceiverInfoLogic) UpdateReceiverInfo(req *types.UpdateReceiverInfoReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderService.UpdateReceiverInfo(l.ctx, &omsclient.UpdateReceiverInfoReq{
		OrderId:               req.Id,
		Status:                req.Status,
		ReceiverName:          req.ReceiverName,
		ReceiverPhone:         req.ReceiverPhone,
		ReceiverPostCode:      req.ReceiverPostCode,
		ReceiverDetailAddress: req.ReceiverDetailAddress,
		ReceiverProvince:      req.ReceiverProvince,
		ReceiverCity:          req.ReceiverCity,
		ReceiverRegion:        req.ReceiverRegion,
		OperateMan:            l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改收货人信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("修改收货人信息失败")
	}

	return res.Success()
}
