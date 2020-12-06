package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogUpdateLogic {
	return MemberLoginLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(req types.UpdateMemberLoginLogReq) (*types.UpdateMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogUpdate(l.ctx, &umsclient.MemberLoginLogUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberLoginLogResp{}, nil
}
