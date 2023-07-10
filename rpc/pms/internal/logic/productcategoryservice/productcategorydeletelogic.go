package productcategoryservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryDeleteLogic {
	return &ProductCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryDeleteLogic) ProductCategoryDelete(in *pmsclient.ProductCategoryDeleteReq) (*pmsclient.ProductCategoryDeleteResp, error) {
	err := l.svcCtx.PmsProductCategoryModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryDeleteResp{}, nil
}
