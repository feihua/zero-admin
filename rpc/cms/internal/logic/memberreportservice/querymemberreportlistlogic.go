package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberReportListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReportListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportListLogic {
	return &QueryMemberReportListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户举报表列表
func (l *QueryMemberReportListLogic) QueryMemberReportList(in *cmsclient.QueryMemberReportListReq) (*cmsclient.QueryMemberReportListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryMemberReportListResp{}, nil
}
