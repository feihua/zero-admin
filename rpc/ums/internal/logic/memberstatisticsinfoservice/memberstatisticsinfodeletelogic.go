package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberStatisticsInfoDeleteLogic 会员统计信息
/*
Author: LiuFeiHua
Date: 2024/5/7 9:21
*/
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

// MemberStatisticsInfoDelete 删除会员统计信息
func (l *MemberStatisticsInfoDeleteLogic) MemberStatisticsInfoDelete(in *umsclient.MemberStatisticsInfoDeleteReq) (*umsclient.MemberStatisticsInfoDeleteResp, error) {
	q := query.UmsGrowthChangeHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberStatisticsInfoDeleteResp{}, nil
}
