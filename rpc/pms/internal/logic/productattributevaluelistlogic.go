package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueListLogic {
	return &ProductAttributeValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueListLogic) ProductAttributeValueList(in *pms.ProductAttributeValueListReq) (*pms.ProductAttributeValueListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductAttributeValueListResp{}, nil
}
