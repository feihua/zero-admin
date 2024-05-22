package integrationchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryListLogic {
	return &IntegrationChangeHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询积分变化历史记录列表
func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(in *umsclient.IntegrationChangeHistoryListReq) (*umsclient.IntegrationChangeHistoryListResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.IntegrationChangeHistoryListResp{}, nil
}
