package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatisticsLoginLogLogic
/*
Author: LiuFeiHua
Date: 2024/01/15 11:51
*/
type StatisticsLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLoginLogLogic {
	return &StatisticsLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// StatisticsLoginLog 统计后台用户登录---(查询当天登录人数（根据IP,统计当前周登录人数（根据IP）,统计当前月登录人数（根据IP）)
func (l *StatisticsLoginLogLogic) StatisticsLoginLog(req *types.StatisticsLoginLogReq) (*types.StatisticsLoginLogResp, error) {
	resp, err := l.svcCtx.LoginLogService.StatisticsLoginLog(l.ctx, &sysclient.StatisticsLoginLogReq{
		Status: "1",
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,统计后台用户登录异常:%+v", req, err.Error())
		return nil, errorx.NewDefaultError("统计后台用户登录失败")
	}

	data := types.StatisticsLoginLogData{
		DayLoginCount:   resp.CurrentDayLoginCount,
		WeekLoginCount:  resp.CurrentWeekLoginCount,
		MonthLoginCount: resp.CurrentMonthLoginCount,
	}
	return &types.StatisticsLoginLogResp{
		Code:    "000000",
		Message: "查询统计后台用户登录",
		Data:    data,
	}, nil
}
