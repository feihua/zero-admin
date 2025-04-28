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

// UpdateOrderReturnReasonLogic 更新退货原因
/*
Author: LiuFeiHua
Date: 2024/6/15 11:43
*/
type UpdateOrderReturnReasonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonLogic {
	return &UpdateOrderReturnReasonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderReturnReason 更新退货原因
func (l *UpdateOrderReturnReasonLogic) UpdateOrderReturnReason(req *types.UpdateOrderReturnReasonReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.UpdateOrderReturnReason(l.ctx, &omsclient.UpdateOrderReturnReasonReq{
		Id:     req.Id,
		Name:   req.Name,   // 退货类型
		Sort:   req.Sort,   // 排序
		Status: req.Status, // 状态：0->不启用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新退货原因失败")
	}

	return res.Success()
}
