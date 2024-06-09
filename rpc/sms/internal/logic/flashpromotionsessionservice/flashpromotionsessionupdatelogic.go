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

// FlashPromotionSessionUpdateLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:47
*/
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

// FlashPromotionSessionUpdate 更新限时购场次
func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(in *smsclient.FlashPromotionSessionUpdateReq) (*smsclient.FlashPromotionSessionUpdateResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotionSession{
		ID:         in.Id,
		Name:       in.Name,
		StartTime:  in.StartTime,
		EndTime:    in.EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionSessionUpdateResp{}, nil
}
