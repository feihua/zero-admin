package log

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryStatisticsLoginLogLogic 统计后台用户登录
/*
Author: LiuFeiHua
Date: 2024/5/29 17:46
*/
type QueryStatisticsLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStatisticsLoginLogLogic {
	return &QueryStatisticsLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryStatisticsLoginLog 统计后台用户登录
func (l *QueryStatisticsLoginLogLogic) QueryStatisticsLoginLog() (resp *types.QueryStatisticsLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return &types.QueryStatisticsLoginLogResp{
		Code:    "000000",
		Message: "查询统计后台用户登录",
		Data: types.QueryStatisticsLoginLogData{
			DayLoginCount:   0,
			WeekLoginCount:  0,
			MonthLoginCount: 0,
		},
	}, nil
}
