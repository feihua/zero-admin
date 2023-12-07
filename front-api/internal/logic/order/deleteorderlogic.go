package order

import (
	"context"
	"encoding/json"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 9:21
*/
type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOrder 用户删除订单
func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrderReq) (resp *types.DeleteOrderResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.OrderService.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		MemberId: memberId,
		OrderId:  req.OrderID,
	})

	if err != nil {
		return nil, err
	}
	return &types.DeleteOrderResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
