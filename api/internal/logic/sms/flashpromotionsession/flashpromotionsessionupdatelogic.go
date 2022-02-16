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

type FlashPromotionSessionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionUpdateLogic {
	return FlashPromotionSessionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(req types.UpdateFlashPromotionSessionReq) (*types.UpdateFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionSessionUpdate(l.ctx, &smsclient.FlashPromotionSessionUpdateReq{
		Id:         req.Id,
		Name:       req.Name,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新限时购场次表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新限时购场次表失败")
	}

	return &types.UpdateFlashPromotionSessionResp{
		Code:    "000000",
		Message: "更新限时购场次表成功",
	}, nil
}
