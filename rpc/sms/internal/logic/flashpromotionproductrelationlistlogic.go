package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationListLogic {
	return &FlashPromotionProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationListLogic) FlashPromotionProductRelationList(in *sms.FlashPromotionProductRelationListReq) (*sms.FlashPromotionProductRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.FlashPromotionProductRelationListResp{}, nil
}
