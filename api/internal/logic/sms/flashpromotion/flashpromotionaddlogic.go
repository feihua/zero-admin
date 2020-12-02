package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionAddLogic {
	return FlashPromotionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionAddLogic) FlashPromotionAdd(req types.AddFlashPromotionReq) (*types.AddFlashPromotionResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddFlashPromotionResp{}, nil
}
