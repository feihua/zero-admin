package cartitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCartItemDetailLogic 查询购物车详情
/*
Author: LiuFeiHua
Date: 2025/07/03 09:32:54
*/
type QueryCartItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCartItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCartItemDetailLogic {
	return &QueryCartItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCartItemDetail 查询购物车详情
func (l *QueryCartItemDetailLogic) QueryCartItemDetail(in *omsclient.QueryCartItemDetailReq) (*omsclient.CartItemData, error) {
	item, err := query.OmsCartItem.WithContext(l.ctx).Where(query.OmsCartItem.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "购物车不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("购物车不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询购物车异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询购物车异常")
	}

	data := &omsclient.CartItemData{
		Id:                item.ID,                                 // 主键ID
		MemberId:          item.MemberID,                           // 会员ID
		ProductId:         item.ProductID,                          // 商品ID
		ProductSkuId:      item.ProductSkuID,                       // 商品SKU ID
		Quantity:          item.Quantity,                           // 购买数量
		Price:             float32(item.Price),                     // 添加到购物车时的价格
		Selected:          item.Selected,                           // 是否选中 0-未选中 1-选中
		ProductName:       item.ProductName,                        // 商品名称
		ProductSubTitle:   item.ProductSubTitle,                    // 商品副标题
		ProductPic:        item.ProductPic,                         // 商品主图URL
		ProductSkuCode:    item.ProductSkuCode,                     // 商品SKU编码
		ProductSn:         item.ProductSn,                          // 商品货号
		ProductBrand:      item.ProductBrand,                       // 商品品牌
		ProductCategoryId: item.ProductCategoryID,                  // 商品分类ID
		ProductAttr:       item.ProductAttr,                        // 商品销售属性JSON
		MemberNickname:    item.MemberNickname,                     // 会员昵称
		Source:            item.Source,                             // 来源 1-PC 2-H5 3-小程序 4-APP
		DeleteStatus:      item.DeleteStatus,                       // 删除状态 0-正常 1-删除
		ExpireTime:        time_util.TimeToStr(item.ExpireTime),    // 过期时间
		CreateTime:        time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime:        time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
