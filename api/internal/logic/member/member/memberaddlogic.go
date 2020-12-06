package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddLogic {
	return MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req types.AddMemberReq) (*types.AddMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberResp{}, nil
}
