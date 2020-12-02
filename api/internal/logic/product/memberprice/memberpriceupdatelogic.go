package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceUpdateLogic {
	return MemberPriceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceUpdateLogic) MemberPriceUpdate(req types.UpdateMemberPriceReq) (*types.UpdateMemberPriceResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateMemberPriceResp{}, nil
}
