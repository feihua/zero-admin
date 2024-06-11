package memberreportservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReportLogic {
	return &UpdateMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户举报表
func (l *UpdateMemberReportLogic) UpdateMemberReport(in *cmsclient.UpdateMemberReportReq) (*cmsclient.UpdateMemberReportResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateMemberReportResp{}, nil
}
