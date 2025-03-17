package product_attribute

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

// ProductAttributeListLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/15 14:31
*/
type ProductAttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeListLogic {
	return &ProductAttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributeList 查询商品属性
func (l *ProductAttributeListLogic) ProductAttributeList(req *types.ListProductAttributeReq) (resp *types.ListProductAttributeResp, err error) {
	attributeList, er := l.svcCtx.ProductAttributeService.QueryProductAttributeList(l.ctx, &pmsclient.QueryProductAttributeListReq{
		PageNum:                    req.Current,
		PageSize:                   req.PageSize,
		Name:                       strings.TrimSpace(req.Name),
		Type:                       req.Type,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
	})

	if er != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询属性列表异常:%s", req, er.Error())
		return nil, errorx.NewDefaultError("查询属性失败")
	}

	var list []*types.ListProductAttributeData

	for _, item := range attributeList.List {
		list = append(list, &types.ListProductAttributeData{
			Id:                         item.Id,                         //
			ProductAttributeCategoryId: item.ProductAttributeCategoryId, // 商品属性分类id
			Name:                       item.Name,                       // 商品属性分类id
			SelectType:                 item.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
			InputType:                  item.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
			InputList:                  item.InputList,                  // 可选值列表，以逗号隔开
			Sort:                       item.Sort,                       // 排序字段：最高的可以单独上传图片
			FilterType:                 item.FilterType,                 // 分类筛选样式：1->普通；1->颜色
			SearchType:                 item.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
			RelatedStatus:              item.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
			HandAddStatus:              item.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
			Type:                       item.Type,                       // 属性的类型；0->规格；1->参数
		})
	}

	return &types.ListProductAttributeResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    attributeList.Total,
		Code:     "000000",
		Message:  "查询属性成功",
	}, nil
}
