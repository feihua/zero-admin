package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionAddLogic {
	return &FlashPromotionSessionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(in *sms.FlashPromotionSessionAddReq) (*sms.FlashPromotionSessionAddResp, error) {
	_, err := l.svcCtx.SmsFlashPromotionSessionModel.Insert(l.ctx, &smsmodel.SmsFlashPromotionSession{
		Name:       in.Name,
		StartTime:  in.StartTime,
		EndTime:    in.EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionSessionAddResp{}, nil
}
