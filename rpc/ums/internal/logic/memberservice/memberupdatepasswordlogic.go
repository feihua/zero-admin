package memberservicelogic

import (
	"context"
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
	member, _ := l.svcCtx.UmsMemberModel.FindOne(l.ctx, in.Id)
	member.Password = in.Password

	err := l.svcCtx.UmsMemberModel.Update(l.ctx, member)
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberUpdateResp{}, nil
}
