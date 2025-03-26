package flashpromotionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionStatusLogic 更新限时购上下线
type UpdateFlashPromotionStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionStatusLogic {
	return &UpdateFlashPromotionStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotionStatus 更新上下线状态
func (l *UpdateFlashPromotionStatusLogic) UpdateFlashPromotionStatus(in *smsclient.UpdateFlashPromotionStatusReq) (*smsclient.UpdateFlashPromotionStatusResp, error) {
	q := query.SmsFlashPromotion
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新上下线状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新上下线状态失败")
	}

	return &smsclient.UpdateFlashPromotionStatusResp{}, nil
}
