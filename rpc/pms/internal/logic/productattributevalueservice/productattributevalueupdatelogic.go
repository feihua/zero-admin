package productattributevalueservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/pmsmodel"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *ProductAttributeValueUpdateLogic) ProductAttributeValueUpdate(in *pmsclient.ProductAttributeValueUpdateReq) (*pmsclient.ProductAttributeValueUpdateResp, error) {
	err := l.svcCtx.PmsProductAttributeValueModel.Update(l.ctx, &pmsmodel.PmsProductAttributeValue{
		Id:                 in.Id,
		ProductId:          in.ProductId,
		ProductAttributeId: in.ProductAttributeId,
		Value:              sql.NullString{String: in.Value},
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueUpdateResp{}, nil
}
