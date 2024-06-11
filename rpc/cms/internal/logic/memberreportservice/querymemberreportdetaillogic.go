package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberReportDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReportDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportDetailLogic {
	return &QueryMemberReportDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户举报表详情
func (l *QueryMemberReportDetailLogic) QueryMemberReportDetail(in *cmsclient.QueryMemberReportDetailReq) (*cmsclient.QueryMemberReportDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryMemberReportDetailResp{}, nil
}
