package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartItemAddLogic {
	return CartItemAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemAddLogic) CartItemAdd(req types.AddCartItemReq) (*types.AddCartItemResp, error) {
	_, err := l.svcCtx.CartItemService.CartItemAdd(l.ctx, &omsclient.CartItemAddReq{
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          req.MemberId,
		Quantity:          req.Quantity,
		Price:             req.Price,
		ProductPic:        req.ProductPic,
		ProductName:       req.ProductName,
		ProductSubTitle:   req.ProductSubTitle,
		ProductSkuCode:    req.ProductSkuCode,
		MemberNickname:    req.MemberNickname,
		DeleteStatus:      req.DeleteStatus,
		ProductCategoryId: req.ProductCategoryId,
		ProductBrand:      req.ProductBrand,
		ProductSn:         req.ProductSn,
		ProductAttr:       req.ProductAttr,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加购物车信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加购物车失败")
	}

	return &types.AddCartItemResp{
		Code:    "000000",
		Message: "添加购物车成功",
	}, nil
}
