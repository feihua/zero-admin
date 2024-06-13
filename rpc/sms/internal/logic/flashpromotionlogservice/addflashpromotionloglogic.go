package flashpromotionlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionLogLogic 添加限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:38
*/
type AddFlashPromotionLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionLogLogic {
	return &AddFlashPromotionLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotionLog 添加限时购通知记录
func (l *AddFlashPromotionLogLogic) AddFlashPromotionLog(in *smsclient.AddFlashPromotionLogReq) (*smsclient.AddFlashPromotionLogResp, error) {
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

	return &smsclient.AddFlashPromotionLogResp{}, nil
}
