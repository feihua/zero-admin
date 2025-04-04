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

// QueryProductCategoryListLogic 查询商品类别列表逻辑处理
/*
Author: LiuFeiHua
Date: 2024/3/7 下午2:58
*/
type QueryProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryListLogic {
	return &QueryProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryList 查询商品类别列表
func (l *QueryProductCategoryListLogic) QueryProductCategoryList(in *pmsclient.QueryProductCategoryListReq) (*pmsclient.QueryProductCategoryListResp, error) {
	q := query.PmsProductCategory.WithContext(l.ctx)

	if in.ParentId != 2000 {
		q = q.Where(query.PmsProductCategory.ParentID.Eq(in.ParentId))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.PmsProductCategory.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品类别列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品类别列表失败")
	}

	var list []*pmsclient.ProductCategoryListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductCategoryListData{
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
		})
	}

	return &pmsclient.QueryProductCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
