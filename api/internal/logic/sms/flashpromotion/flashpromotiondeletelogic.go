package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionDeleteLogic {
	return FlashPromotionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(req types.DeleteFlashPromotionReq) (*types.DeleteFlashPromotionResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteFlashPromotionResp{}, nil
}
