package flashpromotionlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogAddLogic {
	return &FlashPromotionLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogAddLogic) FlashPromotionLogAdd(in *smsclient.FlashPromotionLogAddReq) (*smsclient.FlashPromotionLogAddResp, error) {
	SubscribeTime, _ := time.Parse("2006-01-02 15:04:05", in.SubscribeTime)
	err := query.SmsFlashPromotionLog.WithContext(l.ctx).Create(&model.SmsFlashPromotionLog{
		MemberID:      in.MemberId,
		ProductID:     in.ProductId,
		MemberPhone:   in.MemberPhone,
		ProductName:   in.ProductName,
		SubscribeTime: SubscribeTime,
		SendTime:      time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionLogAddResp{}, nil
}
