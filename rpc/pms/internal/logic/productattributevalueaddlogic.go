package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeValueAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueAddLogic {
	return &ProductAttributeValueAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueAddLogic) ProductAttributeValueAdd(in *pms.ProductAttributeValueAddReq) (*pms.ProductAttributeValueAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductAttributeValueAddResp{}, nil
}
