package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateAttrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateAttrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateAttrLogic {
	return &CartUpdateAttrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartUpdateAttrLogic) CartUpdateAttr(req *types.CartItemUpdateAttrReq) (resp *types.CartItemUpdateResp, err error) {
	_, _ = l.svcCtx.CartItemService.CartItemUpdate(l.ctx, &omsclient.CartItemUpdateReq{
		Id:                req.Id,
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          req.ProductCategoryId,
		Quantity:          req.Quantity,
		Price:             float32(req.Price),
		ProductPic:        req.ProductPic,
		ProductName:       req.ProductName,
		ProductSubTitle:   req.ProductSubTitle,
		ProductSkuCode:    req.ProductSkuCode,
		DeleteStatus:      req.DeleteStatus,
		ProductCategoryId: req.ProductCategoryId,
		ProductBrand:      req.ProductBrand,
		ProductSn:         req.ProductSn,
		ProductAttr:       req.ProductAttr,
	})

	return &types.CartItemUpdateResp{
		Code:    0,
		Message: "操作成功",
	}, nil

	return
}
