package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryAddLogic {
	return &ProductAttributeCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(in *pms.ProductAttributeCategoryAddReq) (*pms.ProductAttributeCategoryAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeCategoryModel.Insert(pmsmodel.PmsProductAttributeCategory{
		Name:           in.Name,
		AttributeCount: in.AttributeCount,
		ParamCount:     in.ParamCount,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeCategoryAddResp{}, nil
}
