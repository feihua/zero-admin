package order_main

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

// DeleteOrderMainLogic 删除订单
/*
Author: LiuFeiHua
Date: 2025/07/01 10:06:43
*/
type DeleteOrderMainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderMainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderMainLogic {
	return &DeleteOrderMainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOrderMain 删除订单
func (l *DeleteOrderMainLogic) DeleteOrderMain(req *types.DeleteOrderMainReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderService.DeleteOrder(l.ctx, &omsclient.DeleteOrderReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除订单失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除订单成功",
	}, nil
}
