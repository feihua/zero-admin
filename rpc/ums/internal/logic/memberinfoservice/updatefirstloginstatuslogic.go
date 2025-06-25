package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFirstLoginStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFirstLoginStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFirstLoginStatusLogic {
	return &UpdateFirstLoginStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFirstLoginStatus 更新会员首次登录状态
func (l *UpdateFirstLoginStatusLogic) UpdateFirstLoginStatus(in *umsclient.UpdateFirstLoginStatusReq) (*umsclient.UpdateMemberInfoResp, error) {
	q := query.UmsMemberInfo

	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).UpdateSimple(q.FirstLoginStatus.Value(2), q.CouponCount.Value(in.CouponCount))
	if err != nil {
		logc.Errorf(l.ctx, "更新会员首次登录状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员首次登录状态失败")
	}

	return &umsclient.UpdateMemberInfoResp{}, nil
}
