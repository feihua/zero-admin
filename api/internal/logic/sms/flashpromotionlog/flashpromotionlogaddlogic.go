package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

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
	_, err := l.svcCtx.Sms.FlashPromotionLogAdd(l.ctx, &smsclient.FlashPromotionLogAddReq{
		MemberId:      req.MemberId,
		ProductId:     req.ProductId,
		MemberPhone:   req.MemberPhone,
		ProductName:   req.ProductName,
		SubscribeTime: req.SubscribeTime,
		SendTime:      req.SendTime,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddFlashPromotionLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
