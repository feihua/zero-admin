package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/ums/umsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressAddLogic {
	return MemberAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(req types.AddMemberAddressReq) (*types.AddMemberAddressResp, error) {
	_, err := l.svcCtx.Ums.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberAddressResp{}, nil
}
