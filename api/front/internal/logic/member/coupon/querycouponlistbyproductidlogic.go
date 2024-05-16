package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponListByProductIdLogic 获取当前商品相关优惠券
/*
Author: LiuFeiHua
Date: 2024/5/16 15:36
*/
type QueryCouponListByProductIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponListByProductIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListByProductIdLogic {
	return &QueryCouponListByProductIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponListByProductId 获取当前商品相关优惠券
// 1.获取指定商品优惠券
// 2.获取指定分类优惠券
// 3.所有优惠券
func (l *QueryCouponListByProductIdLogic) QueryCouponListByProductId(req *types.QueryCouponListByProductIdReq) (resp *types.ListCouponResp, err error) {
	couponList, _ := l.svcCtx.CouponService.CouponFindByProductIdAndProductCategoryId(l.ctx, &smsclient.CouponFindByProductIdAndProductCategoryIdReq{
		ProductId:         req.ProductId,
		ProductCategoryId: req.ProductCategoryId,
	})
	var list []*types.ListCouponData

	for _, item := range couponList.List {
		list = append(list, &types.ListCouponData{
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
		Data:    nil,
		Code:    0,
		Message: "查询优惠券成功",
	}, nil
}
