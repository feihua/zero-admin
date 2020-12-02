package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyDeleteLogic {
	return ReturnApplyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyDeleteLogic) ReturnApplyDelete(req types.DeleteReturnApplyReq) (*types.DeleteReturnApplyResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteReturnApplyResp{}, nil
}
