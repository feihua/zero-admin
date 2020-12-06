package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberUpdateLogic {
	return MemberUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberUpdateLogic) MemberUpdate(req types.UpdateMemberReq) (*types.UpdateMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberUpdate(l.ctx, &umsclient.MemberUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberResp{}, nil
}
