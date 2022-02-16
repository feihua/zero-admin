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

type FlashPromotionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionAddLogic {
	return FlashPromotionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionAddLogic) FlashPromotionAdd(req types.AddFlashPromotionReq) (*types.AddFlashPromotionResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionAdd(l.ctx, &smsclient.FlashPromotionAddReq{
		Title:      req.Title,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加限时购物记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加限时购表失败")
	}

	return &types.AddFlashPromotionResp{
		Code:    "000000",
		Message: "添加限时购表成功",
	}, nil
}
