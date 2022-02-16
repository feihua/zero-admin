package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogDeleteLogic {
	return FlashPromotionLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(req types.DeleteFlashPromotionLogReq) (*types.DeleteFlashPromotionLogResp, error) {
	_, err := l.svcCtx.Sms.FlashPromotionLogDelete(l.ctx, &smsclient.FlashPromotionLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除限时购通知记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除限时购通知记录失败")
	}
	return &types.DeleteFlashPromotionLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
