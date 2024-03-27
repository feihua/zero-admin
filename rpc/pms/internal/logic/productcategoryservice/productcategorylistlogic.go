package productcategoryservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:45
*/
type ProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryListLogic {
	return &ProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryList 查询商品分类列表
func (l *ProductCategoryListLogic) ProductCategoryList(in *pmsclient.ProductCategoryListReq) (*pmsclient.ProductCategoryListResp, error) {
	categoryList, count, err := l.svcCtx.PmsProductCategoryModel.FindAll(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询商品类别列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductCategoryListData
	for _, item := range *categoryList {

		list = append(list, &pmsclient.ProductCategoryListData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			ProductUnit:  item.ProductUnit,
			NavStatus:    item.NavStatus,
			ShowStatus:   item.ShowStatus,
			Sort:         item.Sort,
			Icon:         item.Icon,
			Keywords:     item.Keywords,
			Description:  item.Description.String,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logc.Infof(l.ctx, "查询商品类别列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.ProductCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
