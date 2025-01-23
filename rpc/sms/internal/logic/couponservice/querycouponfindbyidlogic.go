package couponservicelogic

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
		Id:                                coupon.ID,                              //
		Type:                              coupon.Type,                            // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
		Name:                              coupon.Name,                            // 名称
		Platform:                          coupon.Platform,                        // 使用平台：0->全部；1->移动；2->PC
		Count:                             coupon.Count,                           // 数量
		Amount:                            coupon.Amount,                          // 金额
		PerLimit:                          coupon.PerLimit,                        // 每人限领张数
		MinPoint:                          coupon.MinPoint,                        // 使用门槛；0表示无门槛
		StartTime:                         time_util.TimeToStr(coupon.StartTime),  // 开始时间
		EndTime:                           time_util.TimeToStr(coupon.EndTime),    // 结束时间
		UseType:                           coupon.UseType,                         // 使用类型：0->全场通用；1->指定分类；2->指定商品
		Note:                              coupon.Note,                            // 备注
		PublishCount:                      coupon.PublishCount,                    // 发行数量
		UseCount:                          coupon.UseCount,                        // 已使用数量
		ReceiveCount:                      coupon.ReceiveCount,                    // 领取数量
		EnableTime:                        time_util.TimeToStr(coupon.EnableTime), // 可以领取的日期
		Code:                              coupon.Code,                            // 优惠码
		MemberLevel:                       coupon.MemberLevel,                     // 可领取的会员类型：0->无限时
		CouponProductRelationList:         productRelationList,
		CouponProductCategoryRelationList: categoryRelationList,
	}, nil

}
