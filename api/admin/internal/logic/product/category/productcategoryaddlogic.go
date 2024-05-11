package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.ProductCategoryService.ProductCategoryAdd(l.ctx, &pmsclient.ProductCategoryAddReq{
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: 0,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.NavStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加商品分类失败")
	}

	return &types.AddProductCategoryResp{
		Code:    "000000",
		Message: "添加商品分类成功",
	}, nil
}
