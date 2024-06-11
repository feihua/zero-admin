package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberReportLogic {
	return &DeleteMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户举报表
func (l *DeleteMemberReportLogic) DeleteMemberReport(in *cmsclient.DeleteMemberReportReq) (*cmsclient.DeleteMemberReportResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteMemberReportResp{}, nil
}
