package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionSessionLogic 删除限时购场次表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:49
*/
type DeleteFlashPromotionSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionSessionLogic {
	return &DeleteFlashPromotionSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotionSession 删除限时购场次表
func (l *DeleteFlashPromotionSessionLogic) DeleteFlashPromotionSession(in *smsclient.DeleteFlashPromotionSessionReq) (*smsclient.DeleteFlashPromotionSessionResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteFlashPromotionSessionResp{}, nil
}
