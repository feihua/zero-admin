package productcategoryservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAddLogic {
	return &ProductCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAddLogic) ProductCategoryAdd(in *pmsclient.ProductCategoryAddReq) (*pmsclient.ProductCategoryAddResp, error) {
	_, err := l.svcCtx.PmsProductCategoryModel.Insert(l.ctx, &pmsmodel.PmsProductCategory{
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

	return &pmsclient.ProductCategoryAddResp{}, nil
}
