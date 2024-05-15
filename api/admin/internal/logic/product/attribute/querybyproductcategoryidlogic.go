package attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryByproductCategoryIdLogic 根据商品分类的id获取商品属性及属性分类
/*
Author: LiuFeiHua
Date: 2024/5/15 15:09
*/
type QueryByproductCategoryIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryByproductCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByproductCategoryIdLogic {
	return &QueryByproductCategoryIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryByproductCategoryId 根据商品分类的id获取商品属性及属性分类
func (l *QueryByproductCategoryIdLogic) QueryByproductCategoryId(req *types.QueryByproductCategoryIdReq) (resp *types.QueryByproductCategoryIdResp, err error) {
	result, err := l.svcCtx.ProductAttributeService.QueryByproductCategoryId(l.ctx, &pmsclient.QueryByproductCategoryIdReq{ProductCategoryId: req.ProductCategoryId})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,根据商品分类的id获取商品属性及属性分类异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("根据商品分类的id获取商品属性及属性分类失败")
	}

	var list []*types.QueryByproductCategoryIdData

	for _, item := range result.List {
		list = append(list, &types.QueryByproductCategoryIdData{
			AttributeId:                item.AttributeId,
			ProductAttributeCategoryId: item.AttributeCategoryId,
		})
	}

	return &types.QueryByproductCategoryIdResp{
		Data:    list,
		Code:    "000000",
		Message: "根据商品分类的id获取商品属性及属性分类成功",
	}, nil
}
