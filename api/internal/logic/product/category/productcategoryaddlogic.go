package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryAddLogic {
	return ProductCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryAddLogic) ProductCategoryAdd(req types.AddProductCategoryReq) (*types.AddProductCategoryResp, error) {
	_, err := l.svcCtx.Pms.ProductCategoryAdd(l.ctx, &pmsclient.ProductCategoryAddReq{
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: req.ProductCount,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.NavStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加商品分类信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品分类失败")
	}

	return &types.AddProductCategoryResp{
		Code:    "000000",
		Message: "添加商品分类成功",
	}, nil
}
