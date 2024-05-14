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

// FlashPromotionLogUpdateLogic 限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/5/14 10:53
*/
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

// FlashPromotionLogUpdate 更新限时购通知记录
func (l *FlashPromotionLogUpdateLogic) FlashPromotionLogUpdate(in *smsclient.FlashPromotionLogUpdateReq) (*smsclient.FlashPromotionLogUpdateResp, error) {
	SubscribeTime, _ := time.Parse("2006-01-02 15:04:05", in.SubscribeTime)
	SendTime, _ := time.Parse("2006-01-02 15:04:05", in.SendTime)
	q := query.SmsFlashPromotionLog
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotionLog{
		ID:            in.Id,
		MemberID:      in.MemberId,
		ProductID:     in.ProductId,
		MemberPhone:   in.MemberPhone,
		ProductName:   in.ProductName,
		SubscribeTime: SubscribeTime,
		SendTime:      SendTime,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionLogUpdateResp{}, nil
}
