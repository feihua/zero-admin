package attributecategory

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributecategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryListLogic {
	return &ProductAttributecategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributecategoryListLogic) ProductAttributecategoryList(req *types.ListProductAttributecategoryReq) (resp *types.ListProductAttributecategoryResp, err error) {
	attributeList, er := l.svcCtx.ProductAttributeCategoryService.ProductAttributeCategoryList(l.ctx, &pmsclient.ProductAttributeCategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})

	if er != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询属性分类列表异常:%s", req, er.Error())
		return nil, errorx.NewDefaultError("查询属性分类失败")
	}

	var list []*types.ListtProductAttributecategoryData

	for _, item := range attributeList.List {
		list = append(list, &types.ListtProductAttributecategoryData{
			Id:                     item.Id,
			Name:                   item.Name,
			AttributecategoryCount: item.AttributeCount,
			ParamCount:             item.ParamCount,
		})
	}

	return &types.ListProductAttributecategoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    attributeList.Total,
		Code:     "000000",
		Message:  "查询属性分类成功",
	}, nil
}
