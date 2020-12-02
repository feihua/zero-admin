package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionUpdateLogic {
	return FlashPromotionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(req types.UpdateFlashPromotionReq) (*types.UpdateFlashPromotionResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateFlashPromotionResp{}, nil
}
