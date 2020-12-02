package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogDeleteLogic {
	return FlashPromotionLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(req types.DeleteFlashPromotionLogReq) (*types.DeleteFlashPromotionLogResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteFlashPromotionLogResp{}, nil
}
