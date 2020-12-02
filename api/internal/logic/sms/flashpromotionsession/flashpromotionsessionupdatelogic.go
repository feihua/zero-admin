package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionUpdateLogic {
	return FlashPromotionSessionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(req types.UpdateFlashPromotionSessionReq) (*types.UpdateFlashPromotionSessionResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateFlashPromotionSessionResp{}, nil
}
