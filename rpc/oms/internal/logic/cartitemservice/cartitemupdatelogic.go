package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 16:01
*/
type CartItemUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemUpdateLogic {
	return &CartItemUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartItemUpdate 更新购物车商品
// 1.先更加id删除原来的
// 2.添加新的
func (l *CartItemUpdateLogic) CartItemUpdate(in *omsclient.CartItemUpdateReq) (*omsclient.CartItemUpdateResp, error) {
	//1.先更加id删除原来的
	err := l.svcCtx.OmsCartItemModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	//2.添加新的
	_, err = l.svcCtx.OmsCartItemModel.Insert(l.ctx, &omsmodel.OmsCartItem{
		ProductId:         in.ProductId,
		ProductSkuId:      in.ProductSkuId,
		MemberId:          in.MemberId,
		Quantity:          in.Quantity,
		Price:             float64(in.Price),
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductSubTitle:   in.ProductSubTitle,
		ProductSkuCode:    in.ProductSkuCode,
		MemberNickname:    in.MemberNickname,
		CreateDate:        time.Now(),
		ModifyDate:        time.Now(),
		DeleteStatus:      in.DeleteStatus,
		ProductCategoryId: in.ProductCategoryId,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductAttr:       in.ProductAttr,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemUpdateResp{}, nil
}
