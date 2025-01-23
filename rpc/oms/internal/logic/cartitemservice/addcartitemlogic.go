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
	// 1.购物车是否已经存在商品
	item, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId), q.MemberID.Eq(in.MemberId)).First()
	var err error
	if item != nil {
		// 2.如果有,则更新数量
		item.Quantity = item.Quantity + in.Quantity
		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(item.ID)).Update(q.Quantity, item.Quantity)
	} else {
		// 3.插入数据
		_ = q.WithContext(l.ctx).Create(&model.OmsCartItem{
			ProductID:         in.ProductId,         // 商品id
			ProductSkuID:      in.ProductSkuId,      // 商品库存id
			MemberID:          in.MemberId,          // 会员id
			Quantity:          in.Quantity,          // 购买数量
			Price:             in.Price,             // 添加到购物车的价格
			ProductPic:        in.ProductPic,        // 商品主图
			ProductName:       in.ProductName,       // 商品名称
			ProductSubTitle:   in.ProductSubTitle,   // 商品副标题（卖点）
			ProductSkuCode:    in.ProductSkuCode,    // 商品sku条码
			MemberNickname:    in.MemberNickname,    // 会员昵称
			DeleteStatus:      in.DeleteStatus,      // 是否删除
			ProductCategoryID: in.ProductCategoryId, // 商品分类
			ProductBrand:      in.ProductBrand,      // 商品品牌
			ProductSn:         in.ProductSn,         // 货号
			ProductAttr:       in.ProductAttr,       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
		})

	}
	if err != nil {
		return nil, err
	}

	return &omsclient.AddCartItemResp{}, nil
}
