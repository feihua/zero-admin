package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(in *sms.FlashPromotionSessionUpdateReq) (*sms.FlashPromotionSessionUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sms.FlashPromotionSessionUpdateResp{}, nil
}
