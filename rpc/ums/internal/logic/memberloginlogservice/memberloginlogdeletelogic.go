package memberloginlogservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(in *umsclient.MemberLoginLogDeleteReq) (*umsclient.MemberLoginLogDeleteResp, error) {
	err := l.svcCtx.UmsMemberLoginLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLoginLogDeleteResp{}, nil
}
