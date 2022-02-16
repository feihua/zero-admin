package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponListLogic) CouponList(in *sms.CouponListReq) (*sms.CouponListResp, error) {
	all, err := l.svcCtx.SmsCouponModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsCouponModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询优惠券列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.CouponListData
	for _, coupon := range *all {

		list = append(list, &sms.CouponListData{
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

	return &sms.CouponListResp{
		Total: count,
		List:  list,
	}, nil
}
