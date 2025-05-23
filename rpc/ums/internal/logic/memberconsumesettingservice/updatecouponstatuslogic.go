package memberconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCouponStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponStatusLogic {
	return &UpdateCouponStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponStatus 更新是否可以和优惠券同用
func (l *UpdateCouponStatusLogic) UpdateCouponStatus(in *umsclient.UpdateCouponStatusReq) (*umsclient.UpdateMemberConsumeSettingStatusResp, error) {
	q := query.UmsMemberConsumeSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.CouponStatus, in.CouponStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新积分消费设置状态失败")
	}

	return &umsclient.UpdateMemberConsumeSettingStatusResp{}, nil
}
