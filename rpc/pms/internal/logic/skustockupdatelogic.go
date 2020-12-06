package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type SkuStockUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockUpdateLogic {
	return &SkuStockUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockUpdateLogic) SkuStockUpdate(in *pms.SkuStockUpdateReq) (*pms.SkuStockUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.SkuStockUpdateResp{}, nil
}
