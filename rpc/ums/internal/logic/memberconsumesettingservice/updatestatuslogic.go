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

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateStatus 更新积分消费设置状态
func (l *UpdateStatusLogic) UpdateStatus(in *umsclient.UpdateStatusReq) (*umsclient.UpdateMemberConsumeSettingStatusResp, error) {
	q := query.UmsMemberConsumeSetting

	if in.Status == 1 {
		_, _ = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.Status.Eq(in.Status)).Update(q.Status, 0)
	}

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新积分消费设置状态失败")
	}
	return &umsclient.UpdateMemberConsumeSettingStatusResp{}, nil
}
