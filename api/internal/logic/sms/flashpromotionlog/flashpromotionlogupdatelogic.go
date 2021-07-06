package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogUpdateLogic {
	return FlashPromotionLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogUpdateLogic) FlashPromotionLogUpdate(req types.UpdateFlashPromotionLogReq) (*types.UpdateFlashPromotionLogResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionLogUpdate(l.ctx, &smsclient.FlashPromotionLogUpdateReq{
		Id:            req.Id,
		MemberId:      req.MemberId,
		ProductId:     req.ProductId,
		MemberPhone:   req.MemberPhone,
		ProductName:   req.ProductName,
		SubscribeTime: req.SubscribeTime,
		SendTime:      req.SendTime,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新限时购通知记录失败")
	}

	return &types.UpdateFlashPromotionLogResp{
		Code:    "000000",
		Message: "更新限时购通知记录成功",
	}, nil
}
