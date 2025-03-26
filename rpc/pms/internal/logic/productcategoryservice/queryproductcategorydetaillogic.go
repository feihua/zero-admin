package productcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCategoryDetailLogic 查询产品分类详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type QueryProductCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryDetailLogic {
	return &QueryProductCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryDetail 查询产品分类详情
func (l *QueryProductCategoryDetailLogic) QueryProductCategoryDetail(in *pmsclient.QueryProductCategoryDetailReq) (*pmsclient.QueryProductCategoryDetailResp, error) {
	item, err := query.PmsProductCategory.WithContext(l.ctx).Where(query.PmsProductCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品分类详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品分类详情失败")
	}

	data := &pmsclient.QueryProductCategoryDetailResp{
		Id:           item.ID,           //
		ParentId:     item.ParentID,     // 上机分类的编号：0表示一级分类
		Name:         item.Name,         // 商品分类名称
		Level:        item.Level,        // 分类级别：0->1级；1->2级
		ProductCount: item.ProductCount, // 商品数量
		ProductUnit:  item.ProductUnit,  // 商品数量
		NavStatus:    item.NavStatus,    // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus:   item.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         item.Sort,         // 排序
		Icon:         item.Icon,         // 图标
		Keywords:     item.Keywords,     // 关键字
		Description:  item.Description,  // 描述
	}

	return data, nil
}
