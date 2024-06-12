package cart

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartUpdateAttrLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:58
*/
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

// CartUpdateAttr 修改购物车中商品的规格
func (l *CartUpdateAttrLogic) CartUpdateAttr(req *types.CartItemUpdateAttrReq) (resp *types.CartItemUpdateResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.UpdateCartItem(l.ctx, &omsclient.UpdateCartItemReq{
		Id:                req.Id,
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          memberId,
		Quantity:          req.Quantity,
		Price:             req.Price,
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

}
