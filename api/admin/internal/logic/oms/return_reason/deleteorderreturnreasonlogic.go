package return_reason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnReasonLogic 删除退货原因
/*
Author: LiuFeiHua
Date: 2024/6/15 11:41
*/
type DeleteOrderReturnReasonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderReturnReasonLogic {
	return &DeleteOrderReturnReasonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOrderReturnReason 删除退货原因
func (l *DeleteOrderReturnReasonLogic) DeleteOrderReturnReason(req *types.DeleteOrderReturnReasonReq) (resp *types.DeleteOrderReturnReasonResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.DeleteOrderReturnReason(l.ctx, &omsclient.DeleteOrderReturnReasonReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除退货原因异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除退货原因失败")
	}
	return &types.DeleteOrderReturnReasonResp{
		Code:    "000000",
		Message: "删除退货原因成功",
	}, nil
}
