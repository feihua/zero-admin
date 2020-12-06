package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationUpdateLogic {
	return &FlashPromotionProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationUpdateLogic) FlashPromotionProductRelationUpdate(in *sms.FlashPromotionProductRelationUpdateReq) (*sms.FlashPromotionProductRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sms.FlashPromotionProductRelationUpdateResp{}, nil
}
