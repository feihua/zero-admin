package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogUpdateLogic {
	return &FlashPromotionLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogUpdateLogic) FlashPromotionLogUpdate(in *sms.FlashPromotionLogUpdateReq) (*sms.FlashPromotionLogUpdateResp, error) {
	SubscribeTime, _ := time.Parse("2006-01-02 15:04:05", in.SubscribeTime)
	SendTime, _ := time.Parse("2006-01-02 15:04:05", in.SendTime)
	err := l.svcCtx.SmsFlashPromotionLogModel.Update(smsmodel.SmsFlashPromotionLog{
		Id:            in.Id,
		MemberId:      in.MemberId,
		ProductId:     in.ProductId,
		MemberPhone:   in.MemberPhone,
		ProductName:   in.ProductName,
		SubscribeTime: SubscribeTime,
		SendTime:      SendTime,
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionLogUpdateResp{}, nil
}
