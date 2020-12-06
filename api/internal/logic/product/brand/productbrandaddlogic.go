package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandAddLogic {
	return ProductBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandAddLogic) ProductBrandAdd(req types.AddProductBrandReq) (*types.AddProductBrandResp, error) {
	_, err := l.svcCtx.Pms.BrandAdd(l.ctx, &pmsclient.BrandAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddProductBrandResp{}, nil
}
