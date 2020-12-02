package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SkuStockUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockUpdateLogic {
	return SkuStockUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockUpdateLogic) SkuStockUpdate(req types.UpdateSkuStockReq) (*types.UpdateSkuStockResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateSkuStockResp{}, nil
}
