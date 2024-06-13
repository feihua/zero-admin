package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponFindByIdLogic 根据优惠券id查询优惠券
/*
Author: LiuFeiHua
Date: 2024/6/12 17:35
*/
type QueryCouponFindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponFindByIdLogic {
	return &QueryCouponFindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponFindById 根据优惠券id查询优惠券
func (l *QueryCouponFindByIdLogic) QueryCouponFindById(in *smsclient.QueryCouponFindByIdReq) (*smsclient.QueryCouponFindByIdResp, error) {
	coupon, err := query.SmsCoupon.WithContext(l.ctx).Where(query.SmsCoupon.ID.Eq(in.CouponId)).First()

	if err != nil {
		return nil, errors.New("根据优惠券id查询优惠券异常")
	}

	productRelation, err := query.SmsCouponProductRelation.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券与产品关糸列表信息失败,参数：%+v,异常:%s", in, err.Error)
		return nil, err
	}

	var productRelationList []*smsclient.CouponProductRelationData
	for _, item := range productRelation {

		productRelationList = append(productRelationList, &smsclient.CouponProductRelationData{
			Id:          item.ID,
			CouponId:    item.CouponID,
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			ProductSn:   item.ProductSn,
		})
	}

	categoryRelation, err := query.SmsCouponProductCategoryRelation.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券与产品分类关糸列表信息失败,参数：%+v,异常:%s", in, err.Error)
		return nil, err
	}

	var categoryRelationList []*smsclient.CouponProductCategoryRelationData
	for _, item := range categoryRelation {

		categoryRelationList = append(categoryRelationList, &smsclient.CouponProductCategoryRelationData{
			Id:                  item.ID,
			CouponId:            item.CouponID,
			ProductCategoryId:   item.ProductCategoryID,
			ProductCategoryName: item.ProductCategoryName,
			ParentCategoryName:  item.ParentCategoryName,
		})
	}

	return &smsclient.QueryCouponFindByIdResp{
		Id:                                coupon.ID,
		Type:                              coupon.Type,
		Name:                              coupon.Name,
		Platform:                          coupon.Platform,
		Count:                             coupon.Count,
		Amount:                            coupon.Amount,
		PerLimit:                          coupon.PerLimit,
		MinPoint:                          coupon.MinPoint,
		StartTime:                         coupon.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:                           coupon.EndTime.Format("2006-01-02 15:04:05"),
		UseType:                           coupon.UseType,
		Note:                              coupon.Note,
		PublishCount:                      coupon.PublishCount,
		UseCount:                          coupon.UseCount,
		ReceiveCount:                      coupon.ReceiveCount,
		EnableTime:                        coupon.EnableTime.Format("2006-01-02 15:04:05"),
		Code:                              coupon.Code,
		MemberLevel:                       coupon.MemberLevel,
		CouponProductRelationList:         productRelationList,
		CouponProductCategoryRelationList: categoryRelationList,
	}, nil

}
