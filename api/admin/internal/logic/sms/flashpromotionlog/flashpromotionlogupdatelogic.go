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
	_, err := l.svcCtx.FlashPromotionLogService.FlashPromotionLogUpdate(l.ctx, &smsclient.FlashPromotionLogUpdateReq{
		Id:            req.Id,
		MemberId:      req.MemberId,
		ProductId:     req.ProductId,
		MemberPhone:   req.MemberPhone,
		ProductName:   req.ProductName,
		SubscribeTime: req.SubscribeTime,
		SendTime:      req.SendTime,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购通知记录信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新限时购通知记录失败")
	}

	return &types.UpdateFlashPromotionLogResp{
		Code:    "000000",
		Message: "更新限时购通知记录成功",
	}, nil
}
