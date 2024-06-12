package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCartItemLogic 添加购物车表
/*
Author: LiuFeiHua
Date: 2024/6/12 9:40
*/
type AddCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCartItem 添加购物车表
// 1.购物车是否已经存在商品
// 2.如果有,则更新数量
// 3.如果没有,插入数据
func (l *AddCartItemLogic) AddCartItem(in *omsclient.AddCartItemReq) (*omsclient.AddCartItemResp, error) {
	q := query.OmsCartItem
	//1.购物车是否已经存在商品
	item, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId), q.MemberID.Eq(in.MemberId)).First()
	var err error
	if item != nil {
		//2.如果有,则更新数量
		item.Quantity = item.Quantity + in.Quantity
		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(item.ID)).Update(q.Quantity, item.Quantity)
	} else {
		//3.插入数据
		_ = q.WithContext(l.ctx).Create(&model.OmsCartItem{
			ProductID:         in.ProductId,
			ProductSkuID:      in.ProductSkuId,
			MemberID:          in.MemberId,
			Quantity:          in.Quantity,
			Price:             in.Price,
			ProductPic:        in.ProductPic,
			ProductName:       in.ProductName,
			ProductSubTitle:   in.ProductSubTitle,
			ProductSkuCode:    in.ProductSkuCode,
			MemberNickname:    in.MemberNickname,
			DeleteStatus:      in.DeleteStatus,
			ProductCategoryID: in.ProductCategoryId,
			ProductBrand:      in.ProductBrand,
			ProductSn:         in.ProductSn,
			ProductAttr:       in.ProductAttr,
		})

	}
	if err != nil {
		return nil, err
	}

	return &omsclient.AddCartItemResp{}, nil
}
