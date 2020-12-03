package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogDeleteLogic {
	return MemberLoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(req types.DeleteMemberLoginLogReq) (*types.DeleteMemberLoginLogResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteMemberLoginLogResp{}, nil
}
