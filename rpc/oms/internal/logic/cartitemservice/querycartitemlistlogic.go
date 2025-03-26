package cartitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCartItemListLogic 查询购物车列表
/*
Author: LiuFeiHua
Date: 2024/6/12 9:56
*/
type QueryCartItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCartItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCartItemListLogic {
	return &QueryCartItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCartItemList 查询购物车列表
func (l *QueryCartItemListLogic) QueryCartItemList(in *omsclient.QueryCartItemListReq) (*omsclient.QueryCartItemListResp, error) {
	q := query.OmsCartItem
	result, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询购物车列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询购物车列表失败")
	}
	var list []*omsclient.CartItemListData
	for _, item := range result {
		list = append(list, &omsclient.CartItemListData{
			Id:                item.ID,                                 //
			ProductId:         item.ProductID,                          // 商品id
			ProductSkuId:      item.ProductSkuID,                       // 商品库存id
			MemberId:          item.MemberID,                           // 会员id
			Quantity:          item.Quantity,                           // 购买数量
			Price:             item.Price,                              // 添加到购物车的价格
			ProductPic:        item.ProductPic,                         // 商品主图
			ProductName:       item.ProductName,                        // 商品名称
			ProductSubTitle:   item.ProductSubTitle,                    // 商品副标题（卖点）
			ProductSkuCode:    item.ProductSkuCode,                     // 商品sku条码
			MemberNickname:    item.MemberNickname,                     // 会员昵称
			CreateTime:        time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateTime:        time_util.TimeToString(item.UpdateTime), // 更新时间
			DeleteStatus:      item.DeleteStatus,                       // 是否删除
			ProductCategoryId: item.ProductCategoryID,                  // 商品分类
			ProductBrand:      item.ProductBrand,                       // 商品品牌
			ProductSn:         item.ProductSn,                          // 货号
			ProductAttr:       item.ProductAttr,                        // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
		})
	}

	return &omsclient.QueryCartItemListResp{
		Total: 0,
		List:  list,
	}, nil

}
