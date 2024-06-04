package attribute

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
	attributeList, er := l.svcCtx.ProductAttributeService.ProductAttributeList(l.ctx, &pmsclient.ProductAttributeListReq{
		Current:                    req.Current,
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
			Id:                         item.Id,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			SelectType:                 item.SelectType,
			InputType:                  item.InputType,
			InputList:                  item.InputList,
			Sort:                       item.Sort,
			FilterType:                 item.FilterType,
			SearchType:                 item.SearchType,
			RelatedStatus:              item.RelatedStatus,
			HandAddStatus:              item.HandAddStatus,
			Type:                       item.Type,
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
