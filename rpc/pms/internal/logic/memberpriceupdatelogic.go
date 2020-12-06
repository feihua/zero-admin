package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceUpdateLogic {
	return &MemberPriceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceUpdateLogic) MemberPriceUpdate(in *pms.MemberPriceUpdateReq) (*pms.MemberPriceUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.MemberPriceUpdateResp{}, nil
}
