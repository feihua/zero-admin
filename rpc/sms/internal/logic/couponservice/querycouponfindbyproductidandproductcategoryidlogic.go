package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponFindByProductIdAndProductCategoryIdLogic 根据商品Id和分类id查询可用的优惠券(app)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:37
*/
type QueryCouponFindByProductIdAndProductCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponFindByProductIdAndProductCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponFindByProductIdAndProductCategoryIdLogic {
	return &QueryCouponFindByProductIdAndProductCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponFindByProductIdAndProductCategoryId 根据商品Id和分类id查询可用的优惠券(app)
func (l *QueryCouponFindByProductIdAndProductCategoryIdLogic) QueryCouponFindByProductIdAndProductCategoryId(in *smsclient.CouponFindByProductIdAndProductCategoryIdReq) (*smsclient.CouponFindByProductIdAndProductCategoryIdResp, error) {
	sql := `SELECT *
FROM sms_coupon
WHERE use_type = 0
  AND start_time < now()
  AND end_time > now()
UNION
(SELECT c.*
 FROM sms_coupon_product_category_relation cpc
          LEFT JOIN sms_coupon c ON cpc.coupon_id = c.id
 WHERE c.use_type = 1
   AND c.start_time < now()
   AND c.end_time > now()
   AND cpc.product_category_id = ?)
UNION
(SELECT c.*
 FROM sms_coupon_product_relation cp
          LEFT JOIN sms_coupon c ON cp.coupon_id = c.id
 WHERE c.use_type = 2
   AND c.start_time < now()
   AND c.end_time > now()
   AND cp.product_id = ?)`

	var result []model.SmsCoupon
	db := l.svcCtx.DB
	err := db.WithContext(l.ctx).Raw(sql, in.ProductCategoryId, in.ProductId).Scan(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "查询商品优惠券列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.QueryCouponListData
	for _, coupon := range result {

		list = append(list, &smsclient.QueryCouponListData{
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
		})

		logc.Infof(l.ctx, "查询商品优惠券列表信息,参数：%+v,响应：%+v", in, list)
	}

	return &smsclient.CouponFindByProductIdAndProductCategoryIdResp{
		Total: 0,
		List:  list,
	}, nil

}
