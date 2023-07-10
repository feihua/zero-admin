package flashpromotionservicelogic

import (
	"context"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
)

type FlashPromotionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionDeleteLogic {
	return &FlashPromotionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(in *smsclient.FlashPromotionDeleteReq) (*smsclient.FlashPromotionDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionDeleteResp{}, nil
}
