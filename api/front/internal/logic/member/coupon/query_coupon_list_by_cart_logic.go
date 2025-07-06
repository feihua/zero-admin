package coupon

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/api/front/internal/logic/order/cart"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponListByCartLogic 获取登录会员购物车的相关优惠券
/*
Author: LiuFeiHua
Date: 2025/6/19 11:34
*/
type QueryCouponListByCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponListByCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListByCartLogic {
	return &QueryCouponListByCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponListByCart 获取登录会员购物车的相关优惠券
func (l *QueryCouponListByCartLogic) QueryCouponListByCart(req *types.CouponListByCartReq) (resp *types.CouponListByCartResp, err error) {
	// 1.获取购物车信息
	cartPromotionItemList, err := cart.QueryCartListPromotion(nil, l.ctx, l.svcCtx)

	if err != nil {
		return nil, err
	}
	// 获取该用户所有优惠券
	enableList, disableList := QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
	return &types.CouponListByCartResp{
		Data: types.CouponListByCartData{
			EnableList:  enableList,
			DisableList: disableList,
		},
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}

func QueryCouponList(svcCtx *svc.ServiceContext, ctx context.Context, cartPromotionItemList []types.CarItemtPromotionListData) ([]*smsclient.QueryCouponRecordDetailResp, []*smsclient.QueryCouponRecordDetailResp) {
	// memberId, _ := ctx.Value("memberId").(json.Number).Int64()
	// couponList, _ := svcCtx.CouponRecordService.QueryMemberCouponList(ctx, &smsclient.QueryMemberCouponListReq{
	// 	MemberId: memberId,
	// })

	var enableList = make([]*smsclient.QueryCouponRecordDetailResp, 0)
	var disableList = make([]*smsclient.QueryCouponRecordDetailResp, 0)
	// 根据优惠券使用类型来判断优惠券是否可用
	// for _, couponHistoryDetail := range couponList.List {
	// 	useType := couponHistoryDetail.UseType
	// 	minPoint := couponHistoryDetail.CouponListData.MinPoint
	// 	endTime := couponHistoryDetail.CouponListData.EndTime
	// 	nowTime, _ := time.Parse("2006-01-02 15:04:05", endTime)
	// 	productRelationList := couponHistoryDetail.ProductRelationList
	// 	categoryRelationList := couponHistoryDetail.CategoryRelationList
	// 	if useType == 0 {
	// 		// 0->全场通用
	// 		// 判断是否满足优惠起点
	// 		// 计算购物车商品的总价
	// 		var totalAmount int64
	// 		for _, item := range cartPromotionItemList {
	// 			realPrice := item.Price - item.ReduceAmount
	// 			totalAmount = totalAmount + realPrice*int64(item.Quantity)
	// 		}
	// 		if time.Now().Before(nowTime) && totalAmount-minPoint > 0 {
	//
	// 			enableList = append(enableList, couponHistoryDetail)
	// 		} else {
	// 			disableList = append(disableList, couponHistoryDetail)
	// 		}
	// 	} else if useType == 1 {
	// 		// 1->指定分类
	// 		// 计算指定分类商品的总价
	// 		var productCategoryIds = make(map[int64]int64, 0)
	// 		for _, item := range categoryRelationList {
	// 			productCategoryIds[item.ProductCategoryId] = item.ProductCategoryId
	// 		}
	// 		var totalAmount int64
	// 		for _, item := range cartPromotionItemList {
	// 			_, ok := productCategoryIds[item.ProductCategoryId]
	// 			if ok {
	// 				realPrice := item.Price - item.ReduceAmount
	// 				totalAmount = totalAmount + realPrice*int64(item.Quantity)
	// 			}
	// 		}
	// 		if time.Now().Before(nowTime) && totalAmount-minPoint > 0 {
	//
	// 			enableList = append(enableList, couponHistoryDetail)
	// 		} else {
	// 			disableList = append(disableList, couponHistoryDetail)
	// 		}
	// 	} else if useType == 2 {
	// 		// 2->指定商品
	// 		// 计算指定商品的总价
	// 		var productIds = make(map[int64]int64, 0)
	// 		for _, item := range productRelationList {
	// 			productIds[item.ProductId] = item.ProductId
	// 		}
	// 		var totalAmount int64
	// 		for _, item := range cartPromotionItemList {
	// 			_, ok := productIds[item.ProductId]
	// 			if ok {
	// 				realPrice := item.Price - item.ReduceAmount
	// 				totalAmount = totalAmount + realPrice*int64(item.Quantity)
	// 			}
	// 		}
	// 		if time.Now().Before(nowTime) && totalAmount-minPoint > 0 {
	//
	// 			enableList = append(enableList, couponHistoryDetail)
	// 		} else {
	// 			disableList = append(disableList, couponHistoryDetail)
	// 		}
	// 	}
	// }
	enableListStr, _ := sonic.Marshal(enableList)
	disableListStr, _ := sonic.Marshal(disableList)
	logc.Errorf(ctx, "可用的优惠券,参数:%s", enableListStr)
	logc.Errorf(ctx, "不可用的优惠券,参数:%s", disableListStr)
	return enableList, disableList
}
