package logic

import (
	"context"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRegisterLogic {
	return MemberRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRegisterLogic) MemberRegister(req types.MemberRegisterReq) (*types.MemberRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberRegisterResp{}, nil
}
