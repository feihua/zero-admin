package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseListLogic {
	return HomeAdvertiseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(req types.ListHomeAdvertiseReq) (*types.ListHomeAdvertiseResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListHomeAdvertiseResp{}, nil
}
