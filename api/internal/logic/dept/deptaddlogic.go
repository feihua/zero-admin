package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptAddLogic {
	return DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req types.AddDeptReq) (*types.AddDeptResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddDeptResp{}, nil
}
