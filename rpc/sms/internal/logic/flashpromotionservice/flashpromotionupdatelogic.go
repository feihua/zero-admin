package flashpromotionservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
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
	err := l.svcCtx.SmsFlashPromotionModel.Update(l.ctx, &smsmodel.SmsFlashPromotion{
		Id:        in.Id,
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
