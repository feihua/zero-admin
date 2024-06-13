package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponHistoryDetailListLogic 获取该用户所有优惠券
/*
Author: LiuFeiHua
Date: 2023/12/12 10:44
*/
type QueryCouponHistoryDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponHistoryDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryDetailListLogic {
	return &QueryCouponHistoryDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponHistoryDetailList 获取该用户所有优惠券(包括商品和优惠券,商品分类和优惠券的关联关糸)
func (l *QueryCouponHistoryDetailListLogic) QueryCouponHistoryDetailList(in *smsclient.QueryCouponHistoryDetailListReq) (*smsclient.CouponHistoryDetailListResp, error) {
	//1.查询会员领取的所有优惠券记录
	q := query.SmsCouponHistory.WithContext(l.ctx)
	q = q.Where(query.SmsCouponHistory.UseStatus.Eq(0))

	if in.MemberId != 0 {
		q = q.Where(query.SmsCouponHistory.MemberID.Eq(in.MemberId))
	}

	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券使用历史列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list = make([]*smsclient.CouponHistoryDetailListData, 0)
	for _, item := range result {

		//2.查询会员领取的优惠券
		coupon, _ := query.SmsCoupon.WithContext(l.ctx).Where(query.SmsCoupon.ID.Eq(item.CouponID)).First()

		//3.查询商品和优惠券关糸
		relation := query.SmsCouponProductRelation
		couponProductRelationList, _ := relation.WithContext(l.ctx).Where(relation.CouponID.Eq(item.CouponID)).Find()
		var productRelationList = make([]*smsclient.QueryCouponProductRelationListData, 0)
		for _, productRelation := range couponProductRelationList {
			productRelationList = append(productRelationList, &smsclient.QueryCouponProductRelationListData{
				Id:          productRelation.ID,
				CouponId:    productRelation.CouponID,
				ProductId:   productRelation.ProductID,
				ProductName: productRelation.ProductName,
				ProductSn:   productRelation.ProductSn,
			})
		}

		//4.查询商品分类和优惠券关糸
		categoryRelation := query.SmsCouponProductCategoryRelation
		couponProductCategoryRelationList, _ := categoryRelation.WithContext(l.ctx).Where(categoryRelation.CouponID.Eq(item.CouponID)).Find()
		var productCategoryRelationList = make([]*smsclient.QueryCouponProductCategoryRelationListData, 0)
		for _, productCategoryRelation := range couponProductCategoryRelationList {
			productCategoryRelationList = append(productCategoryRelationList, &smsclient.QueryCouponProductCategoryRelationListData{
				Id:                  productCategoryRelation.ID,
				CouponId:            productCategoryRelation.CouponID,
				ProductCategoryId:   productCategoryRelation.ProductCategoryID,
				ProductCategoryName: productCategoryRelation.ProductCategoryName,
				ParentCategoryName:  productCategoryRelation.ParentCategoryName,
			})
		}
		list = append(list, &smsclient.CouponHistoryDetailListData{
			CouponHistoryListData: &smsclient.CouponHistoryListData{
				Id:             item.ID,
				CouponId:       item.CouponID,
				MemberId:       item.MemberID,
				CouponCode:     item.CouponCode,
				MemberNickname: item.MemberNickname,
				GetType:        item.GetType,
				CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
				UseStatus:      item.UseStatus,
				UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
				OrderId:        item.OrderID,
				OrderSn:        item.OrderSn,
			},
			CouponListData: &smsclient.QueryCouponData{
				Id:           coupon.ID,
				Type:         coupon.Type,
				Name:         coupon.Name,
				Platform:     coupon.Platform,
				Count:        coupon.Count,
				Amount:       coupon.Amount,
				PerLimit:     coupon.PerLimit,
				MinPoint:     coupon.MinPoint,
				StartTime:    coupon.StartTime.Format("2006-01-02 15:04:05"),
				EndTime:      coupon.EndTime.Format("2006-01-02 15:04:05"),
				UseType:      coupon.UseType,
				Note:         coupon.Note,
				PublishCount: coupon.PublishCount,
				UseCount:     coupon.UseCount,
				ReceiveCount: coupon.ReceiveCount,
				EnableTime:   coupon.EnableTime.Format("2006-01-02 15:04:05"),
				Code:         coupon.Code,
				MemberLevel:  coupon.MemberLevel,
			},
			ProductRelationList:  productRelationList,
			CategoryRelationList: productCategoryRelationList,
		})
	}

	return &smsclient.CouponHistoryDetailListResp{
		List: list,
	}, nil
}
