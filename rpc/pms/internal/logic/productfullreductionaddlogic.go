package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionAddLogic {
	return &ProductFullReductionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionAddLogic) ProductFullReductionAdd(in *pms.ProductFullReductionAddReq) (*pms.ProductFullReductionAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductFullReductionAddResp{}, nil
}
