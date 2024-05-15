package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponDetailLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 9:40
*/
type CouponDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDetailLogic {
	return &CouponDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponDetail 获取单个优惠券的详细信息
func (l *CouponDetailLogic) CouponDetail(req *types.CouponDetailReq) (resp *types.CouponDetailResp, err error) {
	item, err := l.svcCtx.CouponService.CouponFindById(l.ctx, &smsclient.CouponFindByIdReq{
		CouponId: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,获取单个优惠券的详细信息异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("获取单个优惠券的详细信息失败")
	}

	coupon := buildCoupon(item)

	coupon.ProductRelationList = buildProductRelation(item)

	coupon.ProductCategoryRelationList = buildCategoryRelation(item)

	return &types.CouponDetailResp{
		Data:    coupon,
		Code:    "000000",
		Message: "获取单个优惠券的详细信息成功",
	}, nil
}

// 优惠券基本信息
func buildCoupon(item *couponservice.CouponFindByIdResp) types.ListCouponData {
	coupon := types.ListCouponData{
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
	}
	return coupon
}

// 优惠券绑定的商品分类
func buildCategoryRelation(item *couponservice.CouponFindByIdResp) []*types.ProductCategoryRelation {
	var relations []*types.ProductCategoryRelation
	for _, category := range item.CouponProductCategoryRelationList {
		relations = append(relations, &types.ProductCategoryRelation{
			Id:                  category.Id,
			CouponId:            category.CouponId,
			ProductCategoryId:   category.ProductCategoryId,
			ProductCategoryName: category.ProductCategoryName,
			ParentCategoryName:  category.ParentCategoryName,
		})
	}
	return relations
}

// 优惠券绑定的商品
func buildProductRelation(item *couponservice.CouponFindByIdResp) []*types.ProductRelation {
	var relationList []*types.ProductRelation
	for _, product := range item.CouponProductRelationList {
		relationList = append(relationList, &types.ProductRelation{
			Id:          product.Id,
			CouponId:    product.CouponId,
			ProductId:   product.ProductId,
			ProductName: product.ProductName,
			ProductSn:   product.ProductSn,
		})
	}
	return relationList
}
