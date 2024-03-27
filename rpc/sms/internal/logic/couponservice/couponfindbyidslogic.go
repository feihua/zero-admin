package couponservicelogic

import (
	"context"
	"encoding/json"

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
	all, err := l.svcCtx.SmsCouponModel.FindAllByIds(l.ctx, in.CouponIds)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询优惠券列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponListData
	for _, coupon := range *all {

		list = append(list, &smsclient.CouponListData{
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
		})

		reqStr, _ := json.Marshal(in)
		listStr, _ := json.Marshal(list)
		logx.WithContext(l.ctx).Infof("查询优惠券列表信息,参数：%s,响应：%s", reqStr, listStr)
	}

	return &smsclient.CouponFindByIdsResp{
		Total: 0,
		List:  list,
	}, nil
}
