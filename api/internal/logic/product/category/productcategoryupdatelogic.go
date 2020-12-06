package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryUpdateLogic {
	return ProductCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(req types.UpdateProductCategoryReq) (*types.UpdateProductCategoryResp, error) {
	_, err := l.svcCtx.Pms.ProductCategoryUpdate(l.ctx, &pmsclient.ProductCategoryUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateProductCategoryResp{}, nil
}
