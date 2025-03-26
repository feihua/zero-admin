package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CartItemService.AddCartItem(l.ctx, &omsclient.AddCartItemReq{
		ProductId:         req.ProductId,
		ProductSkuId:      req.ProductSkuId,
		MemberId:          memberId,
		Quantity:          req.Quantity,
		Price:             req.Price,
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

	if err != nil {
		logc.Errorf(l.ctx, "添加商品进购物车失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartItemAddResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
