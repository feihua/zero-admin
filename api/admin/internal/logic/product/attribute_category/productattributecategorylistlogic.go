package attribute_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryListLogic {
	return &ProductAttributeCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributeCategoryListLogic) ProductAttributeCategoryList(req *types.ListProductAttributeCategoryReq) (resp *types.ListProductAttributeCategoryResp, err error) {
	attributeList, er := l.svcCtx.ProductAttributeCategoryService.QueryProductAttributeCategoryList(l.ctx, &pmsclient.QueryProductAttributeCategoryListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Name:     strings.TrimSpace(req.Name),
	})

	if er != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询属性分类列表异常:%s", req, er.Error())
		return nil, errorx.NewDefaultError("查询属性分类失败")
	}

	var list []*types.ListProductAttributeCategoryData

	for _, item := range attributeList.List {
		list = append(list, &types.ListProductAttributeCategoryData{
			Id:                     item.Id,             //
			Name:                   item.Name,           // 商品属性分类名称
			AttributecategoryCount: item.AttributeCount, // 属性数量
			ParamCount:             item.ParamCount,     // 参数数量
		})
	}

	return &types.ListProductAttributeCategoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    attributeList.Total,
		Code:     "000000",
		Message:  "查询属性分类成功",
	}, nil
}
