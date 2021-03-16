package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnResonAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonAddLogic {
	return ReturnResonAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonAddLogic) ReturnResonAdd(req types.AddReturnResonReq) (*types.AddReturnResonResp, error) {
	_, err := l.svcCtx.Oms.OrderReturnReasonAdd(l.ctx, &omsclient.OrderReturnReasonAddReq{
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
