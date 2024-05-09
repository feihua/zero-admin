package cart

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CarItemListLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:36
*/
type CarItemListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarItemListLogic {
	return &CarItemListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CarItemList 获取某个会员的购物车列表
func (l *CarItemListLogic) CarItemList() (resp *types.CartItemListResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	itemListResp, _ := l.svcCtx.CartItemService.CartItemList(l.ctx, &omsclient.CartItemListReq{MemberId: memberId})

	list := make([]types.CartListData, 0)

	for _, item := range itemListResp.List {
		list = append(list, types.CartListData{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.ProductCategoryId,
			Quantity:          item.Quantity,
			Price:             item.Price,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
		})
	}

	return &types.CartItemListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
