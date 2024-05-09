package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionDeleteLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/8 10:16
*/
type FlashPromotionSessionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionDeleteLogic {
	return &FlashPromotionSessionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionSessionDelete 删除限时购场次
func (l *FlashPromotionSessionDeleteLogic) FlashPromotionSessionDelete(in *smsclient.FlashPromotionSessionDeleteReq) (*smsclient.FlashPromotionSessionDeleteResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionSessionDeleteResp{}, nil
}
