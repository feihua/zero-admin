package productattributevalueservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeValueDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueDeleteLogic {
	return &ProductAttributeValueDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueDeleteLogic) ProductAttributeValueDelete(in *pmsclient.ProductAttributeValueDeleteReq) (*pmsclient.ProductAttributeValueDeleteResp, error) {
	err := l.svcCtx.PmsProductAttributeValueModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueDeleteResp{}, nil
}
