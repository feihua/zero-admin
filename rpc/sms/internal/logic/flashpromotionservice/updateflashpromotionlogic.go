package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionLogic 更新限时购表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:44
*/
type UpdateFlashPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionLogic {
	return &UpdateFlashPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotion 更新限时购表
func (l *UpdateFlashPromotionLogic) UpdateFlashPromotion(in *smsclient.UpdateFlashPromotionReq) (*smsclient.UpdateFlashPromotionResp, error) {
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

	return &smsclient.UpdateFlashPromotionResp{}, nil
}
