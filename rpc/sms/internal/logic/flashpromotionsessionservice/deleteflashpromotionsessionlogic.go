package flashpromotionsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionSessionLogic 删除限时购场次
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

// DeleteFlashPromotionSession 删除限时购场次
func (l *DeleteFlashPromotionSessionLogic) DeleteFlashPromotionSession(in *smsclient.DeleteFlashPromotionSessionReq) (*smsclient.DeleteFlashPromotionSessionResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除限时购场次失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除限时购场次失败")
	}

	return &smsclient.DeleteFlashPromotionSessionResp{}, nil
}
