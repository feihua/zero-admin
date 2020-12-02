package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.AddReturnResonResp{}, nil
}
