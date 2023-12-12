package couponhistoryservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

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
func (l *QueryCouponHistoryDetailListLogic) QueryCouponHistoryDetailList(in *smsclient.CouponHistoryDetailListReq) (*smsclient.CouponHistoryDetailListResp, error) {
	//1.查询会员领取的所有优惠券记录
	allCouponHistory, err := l.svcCtx.SmsCouponHistoryModel.FindAll(l.ctx, &smsclient.CouponHistoryListReq{
		Current:   1,
		PageSize:  100,
		CouponId:  0,
		MemberId:  in.MemberId,
		UseStatus: 0,
	})

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询优惠券使用历史列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list = make([]*smsclient.CouponHistoryDetailListData, 0)
	for _, item := range *allCouponHistory {

		//2.查询会员领取的优惠券
		coupon, _ := l.svcCtx.SmsCouponModel.FindOne(l.ctx, item.CouponId)

		//3.查询商品和优惠券关糸
		couponProductRelationList, _ := l.svcCtx.SmsCouponProductRelationModel.FindAllByCouponId(l.ctx, item.CouponId)
		var productRelationList = make([]*smsclient.CouponProductRelationListData, 0)
		for _, productRelation := range *couponProductRelationList {
			productRelationList = append(productRelationList, &smsclient.CouponProductRelationListData{
				Id:          productRelation.Id,
				CouponId:    productRelation.CouponId,
				ProductId:   productRelation.ProductId,
				ProductName: productRelation.ProductName,
				ProductSn:   productRelation.ProductSn,
			})
		}

		//4.查询商品分类和优惠券关糸
		couponProductCategoryRelationList, _ := l.svcCtx.SmsCouponProductCategoryRelationModel.FindAllByCouponId(l.ctx, item.CouponId)
		var productCategoryRelationList = make([]*smsclient.CouponProductCategoryRelationListData, 0)
		for _, productCategoryRelation := range *couponProductCategoryRelationList {
			productCategoryRelationList = append(productCategoryRelationList, &smsclient.CouponProductCategoryRelationListData{
				Id:                  productCategoryRelation.Id,
				CouponId:            productCategoryRelation.CouponId,
				ProductCategoryId:   productCategoryRelation.ProductCategoryId,
				ProductCategoryName: productCategoryRelation.ProductCategoryName,
				ParentCategoryName:  productCategoryRelation.ParentCategoryName,
			})
		}
		list = append(list, &smsclient.CouponHistoryDetailListData{
			CouponHistoryListData: &smsclient.CouponHistoryListData{
				Id:             item.Id,
				CouponId:       item.CouponId,
				MemberId:       item.MemberId,
				CouponCode:     item.CouponCode,
				MemberNickname: item.MemberNickname,
				GetType:        item.GetType,
				CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
				UseStatus:      item.UseStatus,
				UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
				OrderId:        item.OrderId,
				OrderSn:        item.OrderSn,
			},
			CouponListData: &smsclient.CouponListData{
				Id:           coupon.Id,
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
