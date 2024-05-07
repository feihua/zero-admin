package integrationchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationChangeHistoryDeleteLogic 积分变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 9:38
*/
type IntegrationChangeHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryDeleteLogic {
	return &IntegrationChangeHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationChangeHistoryDelete 删除积分变化历史记录
func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(in *umsclient.IntegrationChangeHistoryDeleteReq) (*umsclient.IntegrationChangeHistoryDeleteResp, error) {
	q := query.UmsIntegrationChangeHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationChangeHistoryDeleteResp{}, nil
}
