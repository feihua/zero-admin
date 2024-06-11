package memberreadhistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberReadHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReadHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReadHistoryDetailLogic {
	return &QueryMemberReadHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户商品浏览历史记录详情
func (l *QueryMemberReadHistoryDetailLogic) QueryMemberReadHistoryDetail(in *umsclient.QueryMemberReadHistoryDetailReq) (*umsclient.QueryMemberReadHistoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.QueryMemberReadHistoryDetailResp{}, nil
}
