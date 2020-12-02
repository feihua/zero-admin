package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionListLogic {
	return FlashPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(req types.ListFlashPromotionReq) (*types.ListFlashPromotionResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListFlashPromotionResp{}, nil
}
