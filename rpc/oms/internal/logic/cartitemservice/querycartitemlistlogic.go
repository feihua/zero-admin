package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/logic/common"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCartItemListLogic 查询购物车表列表
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

// QueryCartItemList 查询购物车表列表
func (l *QueryCartItemListLogic) QueryCartItemList(in *omsclient.QueryCartItemListReq) (*omsclient.QueryCartItemListResp, error) {
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
			CreateDate:        item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateDate:        common.TimeToString(item.UpdateTime),
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryID,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
		})
	}

	return &omsclient.QueryCartItemListResp{
		Total: 0,
		List:  list,
	}, nil

}
