package logic

import (
	"context"

	"zero-admin/common/xerr"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type MemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfoLogic {
	return &MemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberInfoLogic) MemberInfo(in *umsclient.MemberInfoReq) (*umsclient.MemberInfoResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberInfoResp{}, nil
}
