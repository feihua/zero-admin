package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.FlashPromotionLogService.FlashPromotionLogAdd(l.ctx, &smsclient.FlashPromotionLogAddReq{
		MemberId:      req.MemberId,
		ProductId:     req.ProductId,
		MemberPhone:   req.MemberPhone,
		ProductName:   req.ProductName,
		SubscribeTime: req.SubscribeTime,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购通知记录信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加限时购通知记录失败")
	}

	return &types.AddFlashPromotionLogResp{
		Code:    "000000",
		Message: "添加限时购通知记录成功",
	}, nil
}
