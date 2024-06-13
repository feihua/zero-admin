package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionSessionLogic 更新限时购场次表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:50
*/
type UpdateFlashPromotionSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionSessionLogic {
	return &UpdateFlashPromotionSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotionSession 更新限时购场次表
func (l *UpdateFlashPromotionSessionLogic) UpdateFlashPromotionSession(in *smsclient.UpdateFlashPromotionSessionReq) (*smsclient.UpdateFlashPromotionSessionResp, error) {
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

	return &smsclient.UpdateFlashPromotionSessionResp{}, nil
}
