package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionListLogic {
	return &ProductFullReductionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionListLogic) ProductFullReductionList(in *pms.ProductFullReductionListReq) (*pms.ProductFullReductionListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductFullReductionListResp{}, nil
}
