package couponservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
)

type CouponUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponUpdateLogic {
	return &CouponUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponUpdateLogic) CouponUpdate(in *smsclient.CouponUpdateReq) (*smsclient.CouponUpdateResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)
	err := l.svcCtx.SmsCouponModel.Update(l.ctx, &smsmodel.SmsCoupon{
		Id:           in.Id,
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Count:        in.Count,
		Amount:       in.Amount,
		PerLimit:     in.PerLimit,
		MinPoint:     in.MinPoint,
		StartTime:    StartTime,
		EndTime:      EndTime,
		UseType:      in.UseType,
		Note:         in.Note,
		PublishCount: in.PublishCount,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   EnableTime,
		Code:         in.Code,
		MemberLevel:  in.MemberLevel,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponUpdateResp{}, nil
}
