package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryListLogic {
	return &ProductAttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryListLogic) ProductAttributeCategoryList(in *pms.ProductAttributeCategoryListReq) (*pms.ProductAttributeCategoryListResp, error) {
	all, err := l.svcCtx.PmsProductAttributeCategoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductAttributeCategoryModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品属性类别列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.ProductAttributeCategoryListData
	for _, item := range *all {

		list = append(list, &pms.ProductAttributeCategoryListData{
			Id:             item.Id,
			Name:           item.Name,
			AttributeCount: item.AttributeCount,
			ParamCount:     item.ParamCount,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品属性类别列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductAttributeCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
