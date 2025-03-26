package integrationchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryIntegrationChangeHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIntegrationChangeHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationChangeHistoryDetailLogic {
	return &QueryIntegrationChangeHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询积分变化历史记录详情
func (l *QueryIntegrationChangeHistoryDetailLogic) QueryIntegrationChangeHistoryDetail(in *umsclient.QueryIntegrationChangeHistoryDetailReq) (*umsclient.QueryIntegrationChangeHistoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryIntegrationChangeHistoryDetailResp{}, nil
}
