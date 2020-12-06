package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogAddLogic {
	return MemberLoginLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(req types.AddMemberLoginLogReq) (*types.AddMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogAdd(l.ctx, &umsclient.MemberLoginLogAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberLoginLogResp{}, nil
}
