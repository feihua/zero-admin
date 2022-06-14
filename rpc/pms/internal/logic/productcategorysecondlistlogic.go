package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/pms"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategorySecondListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategorySecondListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategorySecondListLogic {
	return &ProductCategorySecondListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategorySecondListLogic) ProductCategorySecondList(in *pmsclient.ProductCategorySecondListReq) (*pmsclient.ProductCategoryListResp, error) {
	all, err := l.svcCtx.PmsProductCategoryModel.FindByParentId(in.ParentId)
	count, _ := l.svcCtx.PmsProductCategoryModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品类别列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.ProductCategoryListData
	for _, item := range *all {

		list = append(list, &pms.ProductCategoryListData{
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
			Description:  item.Description,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品类别列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
