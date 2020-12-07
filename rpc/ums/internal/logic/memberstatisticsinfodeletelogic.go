package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(in *ums.MemberStatisticsInfoDeleteReq) (*ums.MemberStatisticsInfoDeleteResp, error) {
	err := l.svcCtx.UmsMemberStatisticsInfoModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberStatisticsInfoDeleteResp{}, nil
}
