package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceAddLogic {
	return MemberPriceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceAddLogic) MemberPriceAdd(req types.AddMemberPriceReq) (*types.AddMemberPriceResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddMemberPriceResp{}, nil
}
