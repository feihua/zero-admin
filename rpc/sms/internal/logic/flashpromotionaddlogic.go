package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionAddLogic {
	return &FlashPromotionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionAddLogic) FlashPromotionAdd(in *sms.FlashPromotionAddReq) (*sms.FlashPromotionAddResp, error) {
	StartDate, _ := time.Parse("2006-01-02 15:04:05", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02 15:04:05", in.EndDate)
	_, err := l.svcCtx.SmsFlashPromotionModel.Insert(l.ctx, &smsmodel.SmsFlashPromotion{
		Title:      in.Title,
		StartDate:  StartDate,
		EndDate:    EndDate,
		Status:     in.Status,
		CreateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionAddResp{}, nil
}
