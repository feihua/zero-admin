package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/oms/omsclient"

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
	_, err := l.svcCtx.Oms.CartItemAdd(l.ctx, &omsclient.CartItemAddReq{
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          req.MemberId,
		Quantity:          req.Quantity,
		Price:             int64(req.Price),
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加购物车信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加购物车失败")
	}

	return &types.AddCartItemResp{
		Code:    "000000",
		Message: "添加购物车成功",
	}, nil
}
