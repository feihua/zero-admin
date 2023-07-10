package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemAddLogic {
	return &CartItemAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemAddLogic) CartItemAdd(req *types.CartItemAddReq) (resp *types.CartItemAddResp, err error) {
	_, _ = l.svcCtx.CartItemService.CartItemAdd(l.ctx, &omsclient.CartItemAddReq{
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          req.ProductCategoryId,
		Quantity:          req.Quantity,
		Price:             float32(req.Price),
		ProductPic:        req.ProductPic,
		ProductName:       req.ProductName,
		ProductSubTitle:   req.ProductSubTitle,
		ProductSkuCode:    req.ProductSkuCode,
		DeleteStatus:      0,
		ProductCategoryId: req.ProductCategoryId,
		ProductBrand:      req.ProductBrand,
		ProductSn:         req.ProductSn,
		ProductAttr:       req.ProductAttr,
	})

	return &types.CartItemAddResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
