package coupon

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/logic/cart"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponListByCartLogic
/*
Author: LiuFeiHua
Date: 2023/12/12 11:35
*/
type CouponListByCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListByCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListByCartLogic {
	return &CouponListByCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponListByCart 获取登录会员购物车的相关优惠券
func (l *CouponListByCartLogic) CouponListByCart(req *types.CouponListByCartReq) (resp *types.CouponListByCartResp, err error) {

	//1.获取购物车信息
	cartPromotionItemList := cart.QueryCartListPromotion(nil, l.ctx, l.svcCtx)

	//获取该用户所有优惠券
	enableList, disableList := QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
	return &types.CouponListByCartResp{
		Data: types.ListCouponData{
			EnableList:  enableList,
			DisableList: disableList,
		},
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}

func QueryCouponList(svcCtx *svc.ServiceContext, ctx context.Context, cartPromotionItemList []types.CarItemtPromotionListData) ([]*smsclient.CouponHistoryDetailListData, []*smsclient.CouponHistoryDetailListData) {
	memberId, _ := ctx.Value("memberId").(json.Number).Int64()
	historyDetailList, _ := svcCtx.CouponHistoryService.QueryCouponHistoryDetailList(ctx, &smsclient.CouponHistoryDetailListReq{
		MemberId: memberId,
	})

	var enableList = make([]*smsclient.CouponHistoryDetailListData, 0)
	var disableList = make([]*smsclient.CouponHistoryDetailListData, 0)
	//根据优惠券使用类型来判断优惠券是否可用
	for _, couponHistoryDetail := range historyDetailList.List {
		useType := couponHistoryDetail.CouponListData.UseType
		minPoint := couponHistoryDetail.CouponListData.MinPoint
		endTime := couponHistoryDetail.CouponListData.EndTime
		nowTime, _ := time.Parse("2006-01-02 15:04:05", endTime)
		productRelationList := couponHistoryDetail.ProductRelationList
		categoryRelationList := couponHistoryDetail.CategoryRelationList
		if useType == 0 {
			//0->全场通用
			//判断是否满足优惠起点
			//计算购物车商品的总价
			var totalAmount float32
			for _, item := range cartPromotionItemList {
				realPrice := item.Price - item.ReduceAmount
				totalAmount = totalAmount + realPrice*float32(item.Quantity)
			}
			if time.Now().Before(nowTime) && totalAmount-float32(minPoint) > 0 {

				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		} else if useType == 1 {
			//1->指定分类
			//计算指定分类商品的总价
			var productCategoryIds = make(map[int64]int64, 0)
			for _, item := range categoryRelationList {
				productCategoryIds[item.ProductCategoryId] = item.ProductCategoryId
			}
			var totalAmount float32
			for _, item := range cartPromotionItemList {
				_, ok := productCategoryIds[item.ProductCategoryId]
				if ok {
					realPrice := item.Price - item.ReduceAmount
					totalAmount = totalAmount + realPrice*float32(item.Quantity)
				}
			}
			if time.Now().Before(nowTime) && totalAmount-float32(minPoint) > 0 {

				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		} else if useType == 2 {
			//2->指定商品
			//计算指定商品的总价
			var productIds = make(map[int64]int64, 0)
			for _, item := range productRelationList {
				productIds[item.ProductId] = item.ProductId
			}
			var totalAmount float32
			for _, item := range cartPromotionItemList {
				_, ok := productIds[item.ProductId]
				if ok {
					realPrice := item.Price - item.ReduceAmount
					totalAmount = totalAmount + realPrice*float32(item.Quantity)
				}
			}
			if time.Now().Before(nowTime) && totalAmount-float32(minPoint) > 0 {

				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		}
	}
	enableListStr, _ := json.Marshal(enableList)
	disableListStr, _ := json.Marshal(disableList)
	logc.Errorf(ctx, "可用的优惠券,参数:%s", enableListStr)
	logc.Errorf(ctx, "不可用的优惠券,参数:%s", disableListStr)
	return enableList, disableList
}
