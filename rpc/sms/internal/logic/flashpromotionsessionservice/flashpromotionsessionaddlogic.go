package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

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

func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(in *smsclient.FlashPromotionSessionAddReq) (*smsclient.FlashPromotionSessionAddResp, error) {
	StartTime, _ := time.Parse("15:04:05", in.StartTime)
	EndTime, _ := time.Parse("15:04:05", in.EndTime)
	err := query.SmsFlashPromotionSession.WithContext(l.ctx).Create(&model.SmsFlashPromotionSession{
		Name:       in.Name,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionSessionAddResp{}, nil
}
