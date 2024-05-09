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

type FlashPromotionSessionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionUpdateLogic {
	return &FlashPromotionSessionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(in *smsclient.FlashPromotionSessionUpdateReq) (*smsclient.FlashPromotionSessionUpdateResp, error) {
	StartTime, _ := time.Parse("15:04:05", in.StartTime)
	EndTime, _ := time.Parse("15:04:05", in.EndTime)
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotionSession{
		ID:         in.Id,
		Name:       in.Name,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionSessionUpdateResp{}, nil
}
