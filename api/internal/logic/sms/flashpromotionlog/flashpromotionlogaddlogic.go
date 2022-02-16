package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Sms.FlashPromotionLogAdd(l.ctx, &smsclient.FlashPromotionLogAddReq{
		MemberId:      req.MemberId,
		ProductId:     req.ProductId,
		MemberPhone:   req.MemberPhone,
		ProductName:   req.ProductName,
		SubscribeTime: req.SubscribeTime,
		SendTime:      req.SendTime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加限时购通知记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加限时购通知记录失败")
	}

	return &types.AddFlashPromotionLogResp{
		Code:    "000000",
		Message: "添加限时购通知记录成功",
	}, nil
}
