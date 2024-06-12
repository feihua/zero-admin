package returnreason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateReturnResonStatusLogic 退货原因
/*
Author: LiuFeiHua
Date: 2024/5/14 11:34
*/
type UpdateReturnResonStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateReturnResonStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReturnResonStatusLogic {
	return &UpdateReturnResonStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateReturnResonStatus 批量更新退货原因状态
func (l *UpdateReturnResonStatusLogic) UpdateReturnResonStatus(req *types.UpdateReturnResonStatusReq) (resp *types.UpdateReturnResonResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.UpdateOrderReturnReasonStatus(l.ctx, &omsclient.UpdateOrderReturnReasonStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量更新退货原因状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量更新退货原因状态失败")
	}

	return &types.UpdateReturnResonResp{
		Code:    "000000",
		Message: "批量更新退货原因状态成功",
	}, nil
}
