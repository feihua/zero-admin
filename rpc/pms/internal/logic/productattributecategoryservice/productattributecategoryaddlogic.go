package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/pmsmodel"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

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

func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(in *pmsclient.ProductAttributeCategoryAddReq) (*pmsclient.ProductAttributeCategoryAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeCategoryModel.Insert(l.ctx, &pmsmodel.PmsProductAttributeCategory{
		Name:           in.Name,
		AttributeCount: 0,
		ParamCount:     0,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeCategoryAddResp{}, nil
}
