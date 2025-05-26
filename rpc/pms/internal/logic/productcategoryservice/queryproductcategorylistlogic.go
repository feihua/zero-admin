package productcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCategoryListLogic 查询产品分类列表
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
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

// QueryProductCategoryList 查询产品分类列表
func (l *QueryProductCategoryListLogic) QueryProductCategoryList(in *pmsclient.QueryProductCategoryListReq) (*pmsclient.QueryProductCategoryListResp, error) {
	productCategory := query.PmsProductCategory
	q := productCategory.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(productCategory.Name.Like("%" + in.Name + "%"))
	}
	if in.NavStatus != 2 {
		q = q.Where(productCategory.NavStatus.Eq(in.NavStatus))
	}
	if len(in.Keywords) > 0 {
		q = q.Where(productCategory.Keywords.Like("%" + in.Keywords + "%"))
	}
	if in.IsEnabled != 2 {
		q = q.Where(productCategory.IsEnabled.Eq(in.IsEnabled))
	}
	if in.ParentId != 1000 {
		q = q.Where(productCategory.ParentID.Eq(in.ParentId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询产品分类列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品分类列表失败")
	}

	var list []*pmsclient.ProductCategoryListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductCategoryListData{
			Id:           item.ID,                                          //
			ParentId:     item.ParentID,                                    // 上级分类的编号：0表示一级分类
			Name:         item.Name,                                        // 商品分类名称
			Level:        item.Level,                                       // 分类级别：0->1级；1->2级
			ProductCount: item.ProductCount,                                // 商品数量
			ProductUnit:  item.ProductUnit,                                 // 商品单位
			NavStatus:    item.NavStatus,                                   // 是否显示在导航栏：0->不显示；1->显示
			Sort:         item.Sort,                                        // 排序
			Icon:         item.Icon,                                        // 图标
			Keywords:     item.Keywords,                                    // 关键字
			Description:  item.Description,                                 // 描述
			IsEnabled:    item.IsEnabled,                                   // 是否启用
			CreateBy:     item.CreateBy,                                    // 创建人ID
			CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
