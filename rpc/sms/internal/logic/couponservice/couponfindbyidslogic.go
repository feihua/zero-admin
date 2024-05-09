package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponFindByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponFindByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponFindByIdsLogic {
	return &CouponFindByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponFindByIds 根据优惠券ids查询优惠券
func (l *CouponFindByIdsLogic) CouponFindByIds(in *smsclient.CouponFindByIdsReq) (*smsclient.CouponFindByIdsResp, error) {

	q := query.SmsCoupon
	result, err := q.WithContext(l.ctx).Where(q.ID.In(in.CouponIds...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponListData
	for _, coupon := range result {

		list = append(list, &smsclient.CouponListData{
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

		logc.Infof(l.ctx, "查询优惠券列表信息,参数：%+v,响应：%+v", in, list)
	}

	return &smsclient.CouponFindByIdsResp{
		Total: 0,
		List:  list,
	}, nil
}
