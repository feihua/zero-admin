package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(in *umsclient.MemberStatisticsInfoAddReq) (*umsclient.MemberStatisticsInfoAddResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberStatisticsInfoAddResp{}, nil
}
