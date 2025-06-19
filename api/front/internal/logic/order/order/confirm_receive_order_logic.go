package order

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfirmReceiveOrderLogic 用户确认收货
/*
Author: LiuFeiHua
Date: 2025/6/20 9:52
*/
type ConfirmReceiveOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmReceiveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmReceiveOrderLogic {
	return &ConfirmReceiveOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfirmReceiveOrder 用户确认收货
func (l *ConfirmReceiveOrderLogic) ConfirmReceiveOrder(req *types.ConfirmReceiveOrderReq) (resp *types.ConfirmReceiveOrderResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderService.OrderConfirm(l.ctx, &omsclient.OrderConfirmReq{
		MemberId: memberId,
		OrderId:  req.OrderId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "用户确认收货失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ConfirmReceiveOrderResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
