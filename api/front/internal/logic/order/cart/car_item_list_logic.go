package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CarItemListLogic 获取某个会员的购物车列表
/*
Author: LiuFeiHua
Date: 2025/6/20 10:22
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
			Id:                item.Id,                //
			ProductId:         item.ProductId,         // 商品id
			ProductSkuId:      item.ProductSkuId,      // 商品库存id
			MemberId:          item.MemberId,          // 会员id
			Quantity:          item.Quantity,          // 购买数量
			Price:             item.Price,             // 添加到购物车的价格
			ProductPic:        item.ProductPic,        // 商品主图
			ProductName:       item.ProductName,       // 商品名称
			ProductSubTitle:   item.ProductSubTitle,   // 商品副标题（卖点）
			ProductSkuCode:    item.ProductSkuCode,    // 商品sku条码
			MemberNickname:    item.MemberNickname,    // 会员昵称
			DeleteStatus:      item.DeleteStatus,      // 是否删除
			ProductCategoryId: item.ProductCategoryId, // 商品分类
			ProductBrand:      item.ProductBrand,      // 商品品牌
			ProductSn:         item.ProductSn,         // 货号
			ProductAttr:       item.ProductAttr,       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]

		})
	}

	return &types.CartItemListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
