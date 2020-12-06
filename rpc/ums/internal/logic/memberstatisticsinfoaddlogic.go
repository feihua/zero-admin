package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoAddLogic {
	return &MemberStatisticsInfoAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(in *ums.MemberStatisticsInfoAddReq) (*ums.MemberStatisticsInfoAddResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberStatisticsInfoAddResp{}, nil
}
