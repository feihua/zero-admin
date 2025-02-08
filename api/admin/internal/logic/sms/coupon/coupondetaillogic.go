package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	item, err := l.svcCtx.CouponService.QueryCouponFindById(l.ctx, &smsclient.QueryCouponFindByIdReq{
		CouponId: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,获取单个优惠券的详细信息异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
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
func buildCoupon(item *couponservice.QueryCouponFindByIdResp) types.ListCouponData {
	coupon := types.ListCouponData{
		Id:           item.Id,           //
		Type:         item.Type,         // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
		Name:         item.Name,         // 名称
		Platform:     item.Platform,     // 使用平台：0->全部；1->移动；2->PC
		Count:        item.Count,        // 数量
		Amount:       item.Amount,       // 金额
		PerLimit:     item.PerLimit,     // 每人限领张数
		MinPoint:     item.MinPoint,     // 使用门槛；0表示无门槛
		StartTime:    item.StartTime,    // 开始时间
		EndTime:      item.EndTime,      // 结束时间
		UseType:      item.UseType,      // 使用类型：0->全场通用；1->指定分类；2->指定商品
		Note:         item.Note,         // 备注
		PublishCount: item.PublishCount, // 发行数量
		UseCount:     item.UseCount,     // 已使用数量
		ReceiveCount: item.ReceiveCount, // 领取数量
		EnableTime:   item.EnableTime,   // 可以领取的日期
		Code:         item.Code,         // 优惠码
		MemberLevel:  item.MemberLevel,  // 可领取的会员类型：0->无限时
	}
	return coupon
}

// 优惠券绑定的商品分类
func buildCategoryRelation(item *couponservice.QueryCouponFindByIdResp) []*types.ProductCategoryRelation {
	var relations []*types.ProductCategoryRelation
	for _, category := range item.CouponProductCategoryRelationList {
		relations = append(relations, &types.ProductCategoryRelation{
			CouponId:            category.CouponId,
			ProductCategoryId:   category.ProductCategoryId,
			ProductCategoryName: category.ProductCategoryName,
			ParentCategoryName:  category.ParentCategoryName,
		})
	}
	return relations
}

// 优惠券绑定的商品
func buildProductRelation(item *couponservice.QueryCouponFindByIdResp) []*types.ProductRelation {
	var relationList []*types.ProductRelation
	for _, product := range item.CouponProductRelationList {
		relationList = append(relationList, &types.ProductRelation{
			CouponId:    product.CouponId,
			ProductId:   product.ProductId,
			ProductName: product.ProductName,
			ProductSn:   product.ProductSn,
		})
	}
	return relationList
}
