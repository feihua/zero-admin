package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
	// 1.查询会员领取的所有优惠券记录
	q := query.SmsCouponHistory.WithContext(l.ctx)
	q = q.Where(query.SmsCouponHistory.UseStatus.Eq(0))

	if in.MemberId != 0 {
		q = q.Where(query.SmsCouponHistory.MemberID.Eq(in.MemberId))
	}

	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "获取该用户所有优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("获取该用户所有优惠券失败")
	}

	var list = make([]*smsclient.CouponHistoryDetailListData, 0)
	for _, item := range result {

		// 2.查询会员领取的优惠券
		coupon, _ := query.SmsCoupon.WithContext(l.ctx).Where(query.SmsCoupon.ID.Eq(item.CouponID)).First()

		// 3.查询商品和优惠券关糸
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

		// 4.查询商品分类和优惠券关糸
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
				Id:             item.ID,                              //
				CouponId:       item.CouponID,                        // 优惠券id
				MemberId:       item.MemberID,                        // 会员id
				CouponCode:     item.CouponCode,                      // 优惠码
				MemberNickname: item.MemberNickname,                  // 领取人昵称
				GetType:        item.GetType,                         // 获取类型：0->后台赠送；1->主动获取
				CreateTime:     time_util.TimeToStr(item.CreateTime), // 领取时间
				UseStatus:      item.UseStatus,                       // 使用状态：0->未使用；1->已使用；2->已过期
				UseTime:        time_util.TimeToStr(item.UseTime),    // 使用时间
				OrderId:        item.OrderID,                         // 订单编号
				OrderSn:        item.OrderSn,                         // 订单号码
			},
			CouponListData: &smsclient.QueryCouponData{
				Id:           coupon.ID,                              //
				Type:         coupon.Type,                            // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
				Name:         coupon.Name,                            // 名称
				Platform:     coupon.Platform,                        // 使用平台：0->全部；1->移动；2->PC
				Count:        coupon.Count,                           // 数量
				Amount:       coupon.Amount,                          // 金额
				PerLimit:     coupon.PerLimit,                        // 每人限领张数
				MinPoint:     coupon.MinPoint,                        // 使用门槛；0表示无门槛
				StartTime:    time_util.TimeToStr(coupon.StartTime),  // 开始时间
				EndTime:      time_util.TimeToStr(coupon.EndTime),    // 结束时间
				UseType:      coupon.UseType,                         // 使用类型：0->全场通用；1->指定分类；2->指定商品
				Note:         coupon.Note,                            // 备注
				PublishCount: coupon.PublishCount,                    // 发行数量
				UseCount:     coupon.UseCount,                        // 已使用数量
				ReceiveCount: coupon.ReceiveCount,                    // 领取数量
				EnableTime:   time_util.TimeToStr(coupon.EnableTime), // 可以领取的日期
				Code:         coupon.Code,                            // 优惠码
				MemberLevel:  coupon.MemberLevel,                     // 可领取的会员类型：0->无限时
			},
			ProductRelationList:  productRelationList,
			CategoryRelationList: productCategoryRelationList,
		})
	}

	return &smsclient.CouponHistoryDetailListResp{
		List: list,
	}, nil
}
