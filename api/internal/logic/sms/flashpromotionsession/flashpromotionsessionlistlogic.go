package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionListLogic {
	return FlashPromotionSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(req types.ListFlashPromotionSessionReq) (*types.ListFlashPromotionSessionResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListFlashPromotionSessionResp{}, nil
}
