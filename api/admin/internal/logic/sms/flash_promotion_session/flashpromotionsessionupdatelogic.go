package flash_promotion_session

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionUpdateLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:44
*/
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

// FlashPromotionSessionUpdate 更新限时购场次
func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(req *types.UpdateFlashPromotionSessionReq) (*types.UpdateFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.FlashPromotionSessionService.UpdateFlashPromotionSession(l.ctx, &smsclient.UpdateFlashPromotionSessionReq{
		Id:        req.Id,
		Name:      req.Name,      // 场次名称
		StartTime: req.StartTime, // 每日开始时间
		EndTime:   req.EndTime,   // 每日结束时间
		Status:    req.Status,    // 启用状态：0->不启用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购场次表信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateFlashPromotionSessionResp{
		Code:    "000000",
		Message: "更新限时购场次表成功",
	}, nil
}
