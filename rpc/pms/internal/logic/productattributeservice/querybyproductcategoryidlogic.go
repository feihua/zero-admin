package productattributeservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryByproductCategoryIdLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/15 15:03
*/
type QueryByproductCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByproductCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByproductCategoryIdLogic {
	return &QueryByproductCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryByproductCategoryId 根据商品分类的id获取商品属性及属性分类
func (l *QueryByproductCategoryIdLogic) QueryByproductCategoryId(in *pmsclient.QueryByproductCategoryIdReq) (*pmsclient.QueryByproductCategoryIdResp, error) {
	db := l.svcCtx.DB
	sql := `SELECT
            t2.id as AttributeId,
            t2.product_attribute_category_id as AttributeCategoryId
        FROM
            pms_product_category_attribute_relation t1
            LEFT JOIN pms_product_attribute t2 ON t2.id = t1.product_attribute_id
        WHERE
            t1.product_category_id = ?`

	var categoryIdData []*pmsclient.QueryByproductCategoryIdData
	err := db.WithContext(l.ctx).Raw(sql, in.ProductCategoryId).Scan(&categoryIdData).Error
	if err != nil {
		logc.Errorf(l.ctx, "根据商品分类的id获取商品属性及属性分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据商品分类的id获取商品属性及属性分类失败")
	}

	return &pmsclient.QueryByproductCategoryIdResp{
		List: categoryIdData,
	}, nil
}
