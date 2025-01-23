package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCartItemLogic 更新购物车表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:05
*/
type UpdateCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCartItem 更新购物车表
// 1.先更加id删除原来的
// 2.添加新的
func (l *UpdateCartItemLogic) UpdateCartItem(in *omsclient.UpdateCartItemReq) (*omsclient.UpdateCartItemResp, error) {
	q := query.OmsCartItem
	// 1.先更加id删除原来的
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()
	if err != nil {
		return nil, err
	}

	// 2.添加新的
	err = q.WithContext(l.ctx).Create(&model.OmsCartItem{
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

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCartItemResp{}, nil
}
