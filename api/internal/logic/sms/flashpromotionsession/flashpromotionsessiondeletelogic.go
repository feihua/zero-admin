package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Sms.FlashPromotionSessionDelete(l.ctx, &smsclient.FlashPromotionSessionDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除限时购场次信息异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除限时购场次信息失败")
	}
	return &types.DeleteFlashPromotionSessionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
