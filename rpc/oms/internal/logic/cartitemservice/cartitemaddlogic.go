package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:46
*/
type CartItemAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemAddLogic {
	return &CartItemAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartItemAdd 添加商品进购物车
func (l *CartItemAddLogic) CartItemAdd(in *omsclient.CartItemAddReq) (*omsclient.CartItemAddResp, error) {
	//1.购物车是否已经存在商品
	item, _ := l.svcCtx.OmsCartItemModel.FindAllByMemberIdAndProduct(l.ctx, in.MemberId, in.ProductId)
	var err error
	if item != nil {
		//2.如果有,则更新数量
		item.Quantity = item.Quantity + in.Quantity
		err = l.svcCtx.OmsCartItemModel.Update(l.ctx, item)
	} else {
		//3.插入数据
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

	}
	if err != nil {
		return nil, err
	}

	return &omsclient.CartItemAddResp{}, nil
}
