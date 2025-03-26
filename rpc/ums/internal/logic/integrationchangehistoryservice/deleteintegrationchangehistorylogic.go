package integrationchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteIntegrationChangeHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteIntegrationChangeHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIntegrationChangeHistoryLogic {
	return &DeleteIntegrationChangeHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除积分变化历史记录
func (l *DeleteIntegrationChangeHistoryLogic) DeleteIntegrationChangeHistory(in *umsclient.DeleteIntegrationChangeHistoryReq) (*umsclient.DeleteIntegrationChangeHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.DeleteIntegrationChangeHistoryResp{}, nil
}
