package productcategoryservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryUpdateLogic {
	return &ProductCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(in *pmsclient.ProductCategoryUpdateReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	err := l.svcCtx.PmsProductCategoryModel.Update(l.ctx, &pmsmodel.PmsProductCategory{
		Id:           in.Id,
		ParentId:     in.ParentId,
		Name:         in.Name,
		Level:        in.Level,
		ProductCount: in.ProductCount,
		ProductUnit:  in.ProductUnit,
		NavStatus:    in.NavStatus,
		ShowStatus:   in.ShowStatus,
		Sort:         in.Sort,
		Icon:         in.Icon,
		Keywords:     in.Keywords,
		Description:  sql.NullString{String: in.Description},
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryUpdateResp{}, nil
}
