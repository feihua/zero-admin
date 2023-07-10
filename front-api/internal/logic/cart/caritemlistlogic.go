package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CarItemListLogic) CarItemList() (resp *types.CartItemListResp, err error) {
	itemListResp, _ := l.svcCtx.CartItemService.CartItemList(l.ctx, &omsclient.CartItemListReq{MemberId: l.ctx.Value("memberId").(int64)})

	var list []types.CartListData

	for _, item := range itemListResp.List {
		list = append(list, types.CartListData{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.ProductCategoryId,
			Quantity:          item.Quantity,
			Price:             float32(item.Price),
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
