package productattributevalueservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ProductAttributeValueAddLogic) ProductAttributeValueAdd(in *pmsclient.ProductAttributeValueAddReq) (*pmsclient.ProductAttributeValueAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeValueModel.Insert(l.ctx, &pmsmodel.PmsProductAttributeValue{
		ProductId:          in.ProductId,
		ProductAttributeId: in.ProductAttributeId,
		Value:              sql.NullString{String: in.Value},
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueAddResp{}, nil
}
