package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemListLogic {
	return &CartItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemListLogic) CartItemList(in *oms.CartItemListReq) (*oms.CartItemListResp, error) {
	all, _ := l.svcCtx.OmsCartItemModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*oms.CartItemListData
	for _, item := range *all {

		list = append(list, &oms.CartItemListData{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.MemberId,
			Quantity:          item.Quantity,
			Price:             int64(item.Price),
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			MemberNickname:    item.MemberNickname,
			CreateDate:        item.CreateDate.Format("2006-01-02 15:04:05"),
			ModifyDate:        item.ModifyDate.Format("2006-01-02 15:04:05"),
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
		})
	}

	fmt.Println(list)
	return &oms.CartItemListResp{
		Total: 10,
		List:  list,
	}, nil
}
