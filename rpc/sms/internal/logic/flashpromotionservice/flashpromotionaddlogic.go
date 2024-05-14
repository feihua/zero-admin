package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionAddLogic 秒杀活动
/*
Author: LiuFeiHua
Date: 2024/5/14 10:54
*/
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

// FlashPromotionAdd 添加秒杀活动
func (l *FlashPromotionAddLogic) FlashPromotionAdd(in *smsclient.FlashPromotionAddReq) (*smsclient.FlashPromotionAddResp, error) {
	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	err := query.SmsFlashPromotion.WithContext(l.ctx).Create(&model.SmsFlashPromotion{
		Title:      in.Title,
		StartDate:  StartDate,
		EndDate:    EndDate,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionAddResp{}, nil
}
