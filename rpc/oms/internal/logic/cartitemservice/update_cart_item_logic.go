package cartitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCartItemLogic 更新购物车
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

// UpdateCartItem 更新购物车
// 1.先删除原来的
// 2.添加新的
func (l *UpdateCartItemLogic) UpdateCartItem(in *omsclient.UpdateCartItemReq) (*omsclient.CartItemResp, error) {
	q := query.OmsCartItem
	// 1.先删除原来的
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expireTime := now.AddDate(0, 0, l.svcCtx.C.Cart.Timeout)
	// 2.添加新的
	err = q.WithContext(l.ctx).Create(&model.OmsCartItem{
		ID:                in.Id,                // 主键ID
		MemberID:          in.MemberId,          // 会员ID
		ProductID:         in.ProductId,         // 商品ID
		ProductSkuID:      in.ProductSkuId,      // 商品SKU ID
		Quantity:          in.Quantity,          // 购买数量
		Price:             float64(in.Price),    // 添加到购物车时的价格
		Selected:          in.Selected,          // 是否选中 0-未选中 1-选中
		ProductName:       in.ProductName,       // 商品名称
		ProductSubTitle:   in.ProductSubTitle,   // 商品副标题
		ProductPic:        in.ProductPic,        // 商品主图URL
		ProductSkuCode:    in.ProductSkuCode,    // 商品SKU编码
		ProductSn:         in.ProductSn,         // 商品货号
		ProductBrand:      in.ProductBrand,      // 商品品牌
		ProductCategoryID: in.ProductCategoryId, // 商品分类ID
		ProductAttr:       in.ProductAttr,       // 商品销售属性JSON
		MemberNickname:    in.MemberNickname,    // 会员昵称
		Source:            in.Source,            // 来源 1-PC 2-H5 3-小程序 4-APP
		ExpireTime:        expireTime,           // 过期时间
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新购物车失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新购物车失败")
	}

	return &omsclient.CartItemResp{}, nil
}
