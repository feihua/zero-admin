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

type MemberAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAuthByAuthKeyLogic {
	return &MemberAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAuthByAuthKeyLogic) MemberAuthByAuthKey(in *umsclient.MemberAuthByAuthKeyReq) (*umsclient.MemberAuthByAuthKeyResp, error) {
	memberAuth, err := l.svcCtx.UmsMemberAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  fail"), "err : %v , in : %+v", err, in)
	}

	var respMemberAuth umsclient.MemberAuth
	_ = copier.Copy(&respMemberAuth, memberAuth)

	return &umsclient.MemberAuthByAuthKeyResp{
		MemberAuth: &respMemberAuth,
	}, nil
}
