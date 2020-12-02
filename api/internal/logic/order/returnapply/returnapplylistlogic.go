package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyListLogic {
	return ReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyListLogic) ReturnApplyList(req types.ListReturnApplyReq) (*types.ListReturnApplyResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListReturnApplyResp{}, nil
}
