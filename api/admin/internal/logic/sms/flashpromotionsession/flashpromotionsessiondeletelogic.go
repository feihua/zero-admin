package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionDeleteLogic {
	return FlashPromotionSessionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionDeleteLogic) FlashPromotionSessionDelete(req types.DeleteFlashPromotionSessionReq) (*types.DeleteFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.FlashPromotionSessionService.FlashPromotionSessionDelete(l.ctx, &smsclient.FlashPromotionSessionDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除限时购场次信息异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除限时购场次信息失败")
	}
	return &types.DeleteFlashPromotionSessionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
