package growthchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GrowthChangeHistoryDeleteLogic 成长值变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 9:39
*/
type GrowthChangeHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryDeleteLogic {
	return &GrowthChangeHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GrowthChangeHistoryDelete 删除成长值变化历史记录
func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(in *umsclient.GrowthChangeHistoryDeleteReq) (*umsclient.GrowthChangeHistoryDeleteResp, error) {
	q := query.UmsGrowthChangeHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.GrowthChangeHistoryDeleteResp{}, nil
}
