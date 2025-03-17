package product_attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeAddLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/15 14:31
*/
type ProductAttributeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeAddLogic {
	return &ProductAttributeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributeAdd 添加商品属性
func (l *ProductAttributeAddLogic) ProductAttributeAdd(req *types.AddProductAttributeReq) (resp *types.AddProductAttributeResp, err error) {
	_, err = l.svcCtx.ProductAttributeService.AddProductAttribute(l.ctx, &pmsclient.AddProductAttributeReq{
		ProductAttributeCategoryId: req.ProductAttributeCategoryId, // 商品属性分类id
		Name:                       req.Name,                       // 商品属性分类id
		SelectType:                 req.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
		InputType:                  req.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
		InputList:                  req.InputList,                  // 可选值列表，以逗号隔开
		Sort:                       req.Sort,                       // 排序字段：最高的可以单独上传图片
		FilterType:                 req.FilterType,                 // 分类筛选样式：1->普通；1->颜色
		SearchType:                 req.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
		RelatedStatus:              req.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
		HandAddStatus:              req.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
		Type:                       req.Type,                       // 属性的类型；0->规格；1->参数
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加属性信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加属性失败")
	}

	return &types.AddProductAttributeResp{
		Code:    "000000",
		Message: "添加属性成功",
	}, nil
}
