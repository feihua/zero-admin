package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
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
	q := query.OmsCartItem
	//1.先更加id删除原来的
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()
	if err != nil {
		return nil, err
	}

	//2.添加新的
	err = q.WithContext(l.ctx).Create(&model.OmsCartItem{
		ProductID:         in.ProductId,
		ProductSkuID:      in.ProductSkuId,
		MemberID:          in.MemberId,
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
		ProductCategoryID: in.ProductCategoryId,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductAttr:       in.ProductAttr,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemUpdateResp{}, nil
}
