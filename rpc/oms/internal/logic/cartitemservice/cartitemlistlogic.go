package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartItemListLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:45
*/
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

// CartItemList 查询购物车列表
func (l *CartItemListLogic) CartItemList(in *omsclient.CartItemListReq) (*omsclient.CartItemListResp, error) {
	q := query.OmsCartItem
	result, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询购物车列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*omsclient.CartItemListData
	for _, item := range result {
		list = append(list, &omsclient.CartItemListData{
			Id:                item.ID,
			ProductId:         item.ProductID,
			ProductSkuId:      item.ProductSkuID,
			MemberId:          item.MemberID,
			Quantity:          item.Quantity,
			Price:             item.Price,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			MemberNickname:    item.MemberNickname,
			CreateDate:        item.CreateDate.Format("2006-01-02 15:04:05"),
			ModifyDate:        item.UpdateDate.Format("2006-01-02 15:04:05"),
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryID,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
		})
	}

	logc.Infof(l.ctx, "查询购物车lis列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.CartItemListResp{
		Total: 0,
		List:  list,
	}, nil
}
