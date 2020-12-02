package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductListLogic {
	return HomeNewProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductListLogic) HomeNewProductList(req types.ListHomeNewProductReq) (*types.ListHomeNewProductResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListHomeNewProductResp{}, nil
}
