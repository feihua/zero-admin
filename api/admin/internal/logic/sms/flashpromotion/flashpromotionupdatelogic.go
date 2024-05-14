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

// FlashPromotionUpdateLogic 秒杀活动
/*
Author: LiuFeiHua
Date: 2024/5/14 10:51
*/
type FlashPromotionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionUpdateLogic {
	return FlashPromotionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionUpdate 更新秒杀活动
func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(req types.UpdateFlashPromotionReq) (*types.UpdateFlashPromotionResp, error) {
	_, err := l.svcCtx.FlashPromotionService.FlashPromotionUpdate(l.ctx, &smsclient.FlashPromotionUpdateReq{
		Id:        req.Id,
		Title:     req.Title,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新限时购表失败")
	}

	return &types.UpdateFlashPromotionResp{
		Code:    "000000",
		Message: "更新限时购表成功",
	}, nil
}
