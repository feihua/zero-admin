package flashpromotionlogservicelogic

import (
	"context"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogDeleteLogic {
	return &FlashPromotionLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(in *smsclient.FlashPromotionLogDeleteReq) (*smsclient.FlashPromotionLogDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionLogDeleteResp{}, nil
}
