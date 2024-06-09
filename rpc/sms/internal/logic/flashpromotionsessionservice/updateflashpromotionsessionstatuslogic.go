package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionSessionStatusLogic 更新上下线状态
type UpdateFlashPromotionSessionStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionSessionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionSessionStatusLogic {
	return &UpdateFlashPromotionSessionStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotionSessionStatus 更新上下线状态
func (l *UpdateFlashPromotionSessionStatusLogic) UpdateFlashPromotionSessionStatus(in *smsclient.UpdateFlashPromotionSessionStatusReq) (*smsclient.UpdateFlashPromotionSessionStatusResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateFlashPromotionSessionStatusResp{}, nil
}
