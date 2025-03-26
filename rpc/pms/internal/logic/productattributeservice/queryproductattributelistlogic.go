package productattributeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeListLogic 查询商品属性参数列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:48
*/
type QueryProductAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeListLogic {
	return &QueryProductAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeList 查询商品属性参数列表
func (l *QueryProductAttributeListLogic) QueryProductAttributeList(in *pmsclient.QueryProductAttributeListReq) (*pmsclient.QueryProductAttributeListResp, error) {
	q := query.PmsProductAttribute.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsProductAttribute.Name.Like("%" + in.Name + "%"))
	}

	if in.Type != 2 {
		q = q.Where(query.PmsProductAttribute.Type.Eq(in.Type))
	}

	if in.ProductAttributeCategoryId != 0 {
		q = q.Where(query.PmsProductAttribute.ProductAttributeCategoryID.Eq(in.ProductAttributeCategoryId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性参数列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品属性参数列表失败")
	}

	var list []*pmsclient.ProductAttributeListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeListData{
			Id:                         item.ID,                         //
			ProductAttributeCategoryId: item.ProductAttributeCategoryID, // 商品属性分类id
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

	return &pmsclient.QueryProductAttributeListResp{
		Total: count,
		List:  list,
	}, nil

}
