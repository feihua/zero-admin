package membertagservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagDeleteLogic {
	return &MemberTagDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(in *umsclient.MemberTagDeleteReq) (*umsclient.MemberTagDeleteResp, error) {
	err := l.svcCtx.UmsMemberTagModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagDeleteResp{}, nil
}
