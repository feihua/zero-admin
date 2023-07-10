package memberstatisticsinfoservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

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
	err := l.svcCtx.UmsMemberStatisticsInfoModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberStatisticsInfoDeleteResp{}, nil
}
