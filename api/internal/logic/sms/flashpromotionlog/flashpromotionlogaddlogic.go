package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogAddLogic {
	return FlashPromotionLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogAddLogic) FlashPromotionLogAdd(req types.AddFlashPromotionLogReq) (*types.AddFlashPromotionLogResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddFlashPromotionLogResp{}, nil
}
