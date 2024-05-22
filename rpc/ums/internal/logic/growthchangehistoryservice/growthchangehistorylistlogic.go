package growthchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryListLogic {
	return &GrowthChangeHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询成长值变化历史记录列表
func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(in *umsclient.GrowthChangeHistoryListReq) (*umsclient.GrowthChangeHistoryListResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.GrowthChangeHistoryListResp{}, nil
}
