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
	_, err := l.svcCtx.FlashPromotionLogService.DeleteFlashPromotionLog(l.ctx, &smsclient.DeleteFlashPromotionLogReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %v,删除限时购通知记录异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除限时购通知记录失败")
	}
	return &types.DeleteFlashPromotionLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
