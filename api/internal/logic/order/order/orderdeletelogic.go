package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDeleteLogic {
	return OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req types.DeleteOrderReq) (*types.DeleteOrderResp, error) {
	_, err := l.svcCtx.Oms.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除订单信息异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除订单信息失败")
	}
	return &types.DeleteOrderResp{
		Code:    "000000",
		Message: "",
	}, nil
}
