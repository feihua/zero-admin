package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductAddLogic {
	return HomeNewProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductAddLogic) HomeNewProductAdd(req types.AddHomeNewProductReq) (*types.AddHomeNewProductResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddHomeNewProductResp{}, nil
}
