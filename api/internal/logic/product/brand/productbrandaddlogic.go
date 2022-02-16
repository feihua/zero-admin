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

type ProductBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandAddLogic {
	return ProductBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandAddLogic) ProductBrandAdd(req types.AddProductBrandReq) (*types.AddProductBrandResp, error) {
	_, err := l.svcCtx.Pms.BrandAdd(l.ctx, &pmsclient.BrandAddReq{
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加商品品牌信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品品牌失败")
	}

	return &types.AddProductBrandResp{
		Code:    "000000",
		Message: "添加商品品牌成功",
	}, nil
}
