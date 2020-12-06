package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeValueUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueUpdateLogic {
	return &ProductAttributeValueUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueUpdateLogic) ProductAttributeValueUpdate(in *pms.ProductAttributeValueUpdateReq) (*pms.ProductAttributeValueUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductAttributeValueUpdateResp{}, nil
}
