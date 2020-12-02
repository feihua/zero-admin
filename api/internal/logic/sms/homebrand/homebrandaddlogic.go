package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandAddLogic {
	return HomeBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandAddLogic) HomeBrandAdd(req types.AddHomeBrandReq) (*types.AddHomeBrandResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddHomeBrandResp{}, nil
}
