package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberReportLogic {
	return &AddMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户举报表
func (l *AddMemberReportLogic) AddMemberReport(in *cmsclient.AddMemberReportReq) (*cmsclient.AddMemberReportResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddMemberReportResp{}, nil
}
