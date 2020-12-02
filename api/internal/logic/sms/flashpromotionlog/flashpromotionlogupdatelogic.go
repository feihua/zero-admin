package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogUpdateLogic {
	return FlashPromotionLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogUpdateLogic) FlashPromotionLogUpdate(req types.UpdateFlashPromotionLogReq) (*types.UpdateFlashPromotionLogResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateFlashPromotionLogResp{}, nil
}
