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

// AddCartItemLogic 添加购物车
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

// AddCartItem 添加购物车
// 1.购物车是否已经存在商品
// 2.如果有,则更新数量
// 3.如果没有,插入数据
func (l *AddCartItemLogic) AddCartItem(in *omsclient.AddCartItemReq) (*omsclient.CartItemResp, error) {
	q := query.OmsCartItem
	// 1.购物车是否已经存在商品
	item, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId), q.MemberID.Eq(in.MemberId)).First()

	now := time.Now()
	expireTime := now.AddDate(0, 0, l.svcCtx.C.Cart.Timeout)
	var err error
	if item != nil {
		// 2.如果有,则更新数量
		item.Quantity = item.Quantity + in.Quantity
		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(item.ID)).Update(q.Quantity, item.Quantity)
	} else {
		// 3.插入数据
		err = q.WithContext(l.ctx).Create(&model.OmsCartItem{
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

	}
	if err != nil {
		logc.Errorf(l.ctx, "添加购物车失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加购物车失败")
	}

	return &omsclient.CartItemResp{}, nil
}
