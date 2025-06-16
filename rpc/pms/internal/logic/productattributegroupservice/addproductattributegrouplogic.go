package productattributegroupservicelogic

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

// AddProductAttributeGroupLogic 添加商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeGroupLogic {
	return &AddProductAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttributeGroup 添加商品属性分组
func (l *AddProductAttributeGroupLogic) AddProductAttributeGroup(in *pmsclient.AddProductAttributeGroupReq) (*pmsclient.AddProductAttributeGroupResp, error) {
	q := query.PmsProductAttributeGroup

	item := &model.PmsProductAttributeGroup{
		CategoryID: in.CategoryId, // 分类ID
		Name:       in.Name,       // 分组名称
		Sort:       in.Sort,       // 排序
		Status:     in.Status,     // 状态：0->禁用；1->启用
		CreateBy:   in.CreateBy,   // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性分组失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品属性分组失败")
	}

	return &pmsclient.AddProductAttributeGroupResp{}, nil
}
