package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionAddLogic {
	return FlashPromotionSessionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(req types.AddFlashPromotionSessionReq) (*types.AddFlashPromotionSessionResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddFlashPromotionSessionResp{}, nil
}
