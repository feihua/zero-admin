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
	var list []*omsclient.CartItemData
	for _, item := range result {
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
		list = append(list, data)
	}

	return &omsclient.QueryCartItemListResp{
		List: list,
	}, nil

}
