package flashpromotionlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionLogLogic 删除限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:39
*/
type DeleteFlashPromotionLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionLogLogic {
	return &DeleteFlashPromotionLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotionLog 删除限时购通知记录
func (l *DeleteFlashPromotionLogLogic) DeleteFlashPromotionLog(in *smsclient.DeleteFlashPromotionLogReq) (*smsclient.DeleteFlashPromotionLogResp, error) {
	q := query.SmsFlashPromotionLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteFlashPromotionLogResp{}, nil
}
