package flashpromotionsession

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionAddLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:43
*/
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

// FlashPromotionSessionAdd 添加限时购场次
func (l *FlashPromotionSessionAddLogic) FlashPromotionSessionAdd(req *types.AddFlashPromotionSessionReq) (*types.AddFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.FlashPromotionSessionService.FlashPromotionSessionAdd(l.ctx, &smsclient.FlashPromotionSessionAddReq{
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购场次信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加限时购场次表失败")
	}

	return &types.AddFlashPromotionSessionResp{
		Code:    "000000",
		Message: "添加限时购场次表成功",
	}, nil
}
