package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

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
	_, err := l.svcCtx.Sms.FlashPromotionSessionAdd(l.ctx, &smsclient.FlashPromotionSessionAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddFlashPromotionSessionResp{}, nil
}
