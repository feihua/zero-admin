package flashpromotionlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionLogDeleteLogic 限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/5/8 10:12
*/
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

// FlashPromotionLogDelete 删除限时购通知记录
func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(in *smsclient.FlashPromotionLogDeleteReq) (*smsclient.FlashPromotionLogDeleteResp, error) {
	q := query.SmsFlashPromotionLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionLogDeleteResp{}, nil
}
