package integrationchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryAddLogic {
	return &IntegrationChangeHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加积分变化历史记录
func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(in *umsclient.IntegrationChangeHistoryAddReq) (*umsclient.IntegrationChangeHistoryAddResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.IntegrationChangeHistoryAddResp{}, nil
}
