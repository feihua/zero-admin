package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberDeleteLogic {
	return MemberDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberDeleteLogic) MemberDelete(req types.DeleteMemberReq) (*types.DeleteMemberResp, error) {
	_, _ = l.svcCtx.Ums.MemberDelete(l.ctx, &umsclient.MemberDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberResp{}, nil
}
