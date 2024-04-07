package order

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

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
		OrderId:  req.OrderId,
	})

	if err != nil {
		return nil, err
	}
	return &types.DeleteOrderResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
