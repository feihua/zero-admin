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

func (l *FlashPromotionUpdateLogic) FlashPromotionUpdate(req types.UpdateFlashPromotionReq) (*types.UpdateFlashPromotionResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionUpdate(l.ctx, &smsclient.FlashPromotionUpdateReq{
		Id:         req.Id,
		Title:      req.Title,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新限时购信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新限时购表失败")
	}

	return &types.UpdateFlashPromotionResp{
		Code:    "000000",
		Message: "更新限时购表成功",
	}, nil
}
