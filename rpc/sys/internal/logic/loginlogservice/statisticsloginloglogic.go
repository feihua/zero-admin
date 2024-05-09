package loginlogservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatisticsLoginLogLogic
/*
Author: LiuFeiHua
Date: 2024/01/15 11:55
*/
type StatisticsLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLoginLogLogic {
	return &StatisticsLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StatisticsLoginLog 统计后台用户登录---(查询当天登录人数（根据IP）,统计当前周登录人数（根据IP）,统计当前月登录人数（根据IP）)
func (l *StatisticsLoginLogLogic) StatisticsLoginLog(in *sysclient.StatisticsLoginLogReq) (*sysclient.StatisticsLoginLogResp, error) {
	//查询当天登录人数（根据IP）
	//dayLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 1)
	//统计当前周登录人数（根据IP）
	//weekLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 2)
	//统计当前月登录人数（根据IP）
	//monthLoginCount, _ := l.svcCtx.LoginLogModel.StatisticsLoginLog(l.ctx, 3)
	return &sysclient.StatisticsLoginLogResp{
		CurrentDayLoginCount:   0,
		CurrentWeekLoginCount:  0,
		CurrentMonthLoginCount: 0,
	}, nil
}
