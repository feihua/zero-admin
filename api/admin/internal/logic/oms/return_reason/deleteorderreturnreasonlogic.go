package return_reason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnReasonLogic 删除退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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
func (l *DeleteOrderReturnReasonLogic) DeleteOrderReturnReason(req *types.DeleteOrderReturnReasonReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.DeleteOrderReturnReason(l.ctx, &omsclient.DeleteOrderReturnReasonReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除退货原因失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除退货原因成功",
	}, nil
}
