package logic

import (
	"context"

	"zero-admin/common/xerr"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
)

type MemberAuthByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAuthByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAuthByMemberIdLogic {
	return &MemberAuthByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAuthByMemberIdLogic) MemberAuthByMemberId(in *umsclient.MemberAuthByMemberIdReq) (*umsclient.MemberAuthyMemberIdResp, error) {
	memberAuth, err := l.svcCtx.UmsMemberAuthModel.FindOneByMemberIdAuthType(l.ctx, in.MemberId, in.AuthType)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  fail"), "err : %v , in : %+v", err, in)
	}

	var respMemberAuth umsclient.MemberAuth
	_ = copier.Copy(&respMemberAuth, memberAuth)

	return &umsclient.MemberAuthyMemberIdResp{
		MemberAuth: &respMemberAuth,
	}, nil
}
