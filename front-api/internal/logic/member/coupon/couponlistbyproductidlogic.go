package coupon

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponListByProductIdLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:34
*/
type CouponListByProductIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListByProductIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListByProductIdLogic {
	return &CouponListByProductIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponListByProductId 获取当前商品相关优惠券
// 1.获取指定商品优惠券
// 2.获取指定分类优惠券
// 3.所有优惠券
func (l *CouponListByProductIdLogic) CouponListByProductId(req *types.ListCouponReq) (resp *types.ListCouponResp, err error) {
	var couponIds []int64
	//1.获取指定商品优惠券
	productRelationList, _ := l.svcCtx.CouponProductRelationService.CouponProductRelationList(l.ctx, &smsclient.CouponProductRelationListReq{
		ProductId: req.ProductId,
	})

	for _, i := range productRelationList.List {
		couponIds = append(couponIds, i.CouponId)
	}

	//2.获取指定分类优惠券
	var ids []int64
	ids = append(ids, req.ProductId)
	productListByIds, _ := l.svcCtx.ProductService.ProductListByIds(l.ctx, &pmsclient.ProductByIdsReq{Ids: ids})
	productCategoryId := productListByIds.List[0].ProductCategoryId

	productCategoryRelationList, _ := l.svcCtx.CouponProductCategoryRelationService.CouponProductCategoryRelationList(l.ctx, &smsclient.CouponProductCategoryRelationListReq{
		ProductCategoryId: productCategoryId,
	})

	for _, i := range productCategoryRelationList.List {
		couponIds = append(couponIds, i.CouponId)
	}

	//3.所有优惠券
	couponFindByIds, _ := l.svcCtx.CouponService.CouponFindByIds(l.ctx, &smsclient.CouponFindByIdsReq{CouponIds: couponIds})

	var list []*types.ListtCouponData

	for _, item := range couponFindByIds.List {
		list = append(list, &types.ListtCouponData{
			Id:           item.Id,
			Type:         item.Type,
			Name:         item.Name,
			Platform:     item.Platform,
			Count:        item.Count,
			Amount:       item.Amount,
			PerLimit:     item.PerLimit,
			MinPoint:     item.MinPoint,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			UseType:      item.UseType,
			Note:         item.Note,
			PublishCount: item.PublishCount,
			UseCount:     item.UseCount,
			ReceiveCount: item.ReceiveCount,
			EnableTime:   item.EnableTime,
			Code:         item.Code,
			MemberLevel:  item.MemberLevel,
		})
	}

	return &types.ListCouponResp{
		Data:    list,
		Code:    0,
		Message: "查询优惠券成功",
	}, nil
}
