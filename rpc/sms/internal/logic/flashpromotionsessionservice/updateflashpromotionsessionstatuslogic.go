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
		logc.Errorf(l.ctx, "更新上下线状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新上下线状态失败")
	}

	return &smsclient.UpdateFlashPromotionSessionStatusResp{}, nil
}
