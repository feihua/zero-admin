package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionUpdateLogic {
	return &ProductFullReductionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionUpdateLogic) ProductFullReductionUpdate(in *pms.ProductFullReductionUpdateReq) (*pms.ProductFullReductionUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductFullReductionUpdateResp{}, nil
}
