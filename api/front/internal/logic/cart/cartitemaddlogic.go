package cart

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:46
*/
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

// CartItemAdd 添加商品进购物车
func (l *CartItemAddLogic) CartItemAdd(req *types.CartItemAddReq) (resp *types.CartItemAddResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.CartItemAdd(l.ctx, &omsclient.CartItemAddReq{
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          memberId,
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
