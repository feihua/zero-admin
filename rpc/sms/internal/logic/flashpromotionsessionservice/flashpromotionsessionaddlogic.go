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

// FlashPromotionSessionAddLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:46
*/
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

// FlashPromotionSessionAdd 添加限时购场次
func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(in *smsclient.FlashPromotionSessionAddReq) (*smsclient.FlashPromotionSessionAddResp, error) {
	err := query.SmsFlashPromotionSession.WithContext(l.ctx).Create(&model.SmsFlashPromotionSession{
		Name:       in.Name,
		StartTime:  in.StartTime,
		EndTime:    in.EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionSessionAddResp{}, nil
}
