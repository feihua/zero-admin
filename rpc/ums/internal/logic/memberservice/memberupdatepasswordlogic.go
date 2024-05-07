package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberUpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdatePasswordLogic {
	return &MemberUpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberUpdatePasswordLogic) MemberUpdatePassword(in *umsclient.MemberUpdatePasswordReq) (*umsclient.MemberUpdateResp, error) {
	q := query.UmsMember
	member, _ := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	member.Password = in.Password

	_, err := q.WithContext(l.ctx).Updates(member)
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberUpdateResp{}, nil
}
