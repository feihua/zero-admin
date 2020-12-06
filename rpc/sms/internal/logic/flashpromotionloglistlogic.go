package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogListLogic {
	return &FlashPromotionLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogListLogic) FlashPromotionLogList(in *sms.FlashPromotionLogListReq) (*sms.FlashPromotionLogListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.FlashPromotionLogListResp{}, nil
}
