package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoDeleteLogic {
	return &MemberStatisticsInfoDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(in *umsclient.MemberStatisticsInfoDeleteReq) (*umsclient.MemberStatisticsInfoDeleteResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsMemberStatisticsInfoModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &umsclient.MemberStatisticsInfoDeleteResp{}, nil
}
