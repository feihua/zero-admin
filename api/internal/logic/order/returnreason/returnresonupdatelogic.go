package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnResonUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonUpdateLogic {
	return ReturnResonUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonUpdateLogic) ReturnResonUpdate(req types.UpdateReturnResonReq) (*types.UpdateReturnResonResp, error) {
	_, err := l.svcCtx.Oms.OrderReturnReasonUpdate(l.ctx, &omsclient.OrderReturnReasonUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateReturnResonResp{}, nil
}
