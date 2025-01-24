package productattributecategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeCategoryLogic 添加产品属性分类
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type AddProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeCategoryLogic {
	return &AddProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttributeCategory 添加产品属性分类
func (l *AddProductAttributeCategoryLogic) AddProductAttributeCategory(in *pmsclient.AddProductAttributeCategoryReq) (*pmsclient.AddProductAttributeCategoryResp, error) {
	q := query.PmsProductAttributeCategory

	item := &model.PmsProductAttributeCategory{
		Name:           in.Name, // 商品属性分类名称
		AttributeCount: 0,       // 属性数量
		ParamCount:     0,       // 参数数量
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加产品属性分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加产品属性分类失败")
	}

	logc.Infof(l.ctx, "添加产品属性分类成功,参数：%+v", in)
	return &pmsclient.AddProductAttributeCategoryResp{}, nil
}
