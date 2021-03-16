package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionDeleteLogic {
	return FlashPromotionSessionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionDeleteLogic) FlashPromotionSessionDelete(req types.DeleteFlashPromotionSessionReq) (*types.DeleteFlashPromotionSessionResp, error) {
	_, _ = l.svcCtx.Sms.FlashPromotionSessionDelete(l.ctx, &smsclient.FlashPromotionSessionDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteFlashPromotionSessionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
