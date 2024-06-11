package growthchangehistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGrowthChangeHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGrowthChangeHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGrowthChangeHistoryDetailLogic {
	return &QueryGrowthChangeHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询成长值变化历史记录表详情
func (l *QueryGrowthChangeHistoryDetailLogic) QueryGrowthChangeHistoryDetail(in *umsclient.QueryGrowthChangeHistoryDetailReq) (*umsclient.QueryGrowthChangeHistoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryGrowthChangeHistoryDetailResp{}, nil
}
