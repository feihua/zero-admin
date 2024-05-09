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

type FlashPromotionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionUpdateLogic {
	return &FlashPromotionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(in *smsclient.FlashPromotionUpdateReq) (*smsclient.FlashPromotionUpdateResp, error) {
	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	q := query.SmsFlashPromotion
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotion{
		ID:        in.Id,
		Title:     in.Title,
		StartDate: StartDate,
		EndDate:   EndDate,
		Status:    in.Status,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionUpdateResp{}, nil
}
