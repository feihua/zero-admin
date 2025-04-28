package return_reason

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

// UpdateOrderReturnReasonStatusLogic 批量更新退货原因状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:44
*/
type UpdateOrderReturnReasonStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderReturnReasonStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonStatusLogic {
	return &UpdateOrderReturnReasonStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderReturnReasonStatus 批量更新退货原因状态
func (l *UpdateOrderReturnReasonStatusLogic) UpdateOrderReturnReasonStatus(req *types.UpdateOrderReturnReasonStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.UpdateOrderReturnReasonStatus(l.ctx, &omsclient.UpdateOrderReturnReasonStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量更新退货原因状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量更新退货原因状态失败")
	}

	return res.Success()
}
