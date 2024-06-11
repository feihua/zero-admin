package growthchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGrowthChangeHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGrowthChangeHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGrowthChangeHistoryLogic {
	return &DeleteGrowthChangeHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除成长值变化历史记录表
func (l *DeleteGrowthChangeHistoryLogic) DeleteGrowthChangeHistory(in *umsclient.DeleteGrowthChangeHistoryReq) (*umsclient.DeleteGrowthChangeHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.DeleteGrowthChangeHistoryResp{}, nil
}
