package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyUpdateLogic {
	return ReturnApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyUpdateLogic) ReturnApplyUpdate(req types.UpdateReturnApplyReq) (*types.UpdateReturnApplyResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateReturnApplyResp{}, nil
}
