package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	itemListResp, err := l.svcCtx.CartItemService.QueryCartItemList(l.ctx, &omsclient.QueryCartItemListReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "获取某个会员的购物车列表失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	list := make([]types.CartListData, 0)

	for _, item := range itemListResp.List {
		list = append(list, types.CartListData{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.MemberId,
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
