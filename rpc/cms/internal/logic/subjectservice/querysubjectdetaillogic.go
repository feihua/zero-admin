package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectDetailLogic {
	return &QuerySubjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题表详情
func (l *QuerySubjectDetailLogic) QuerySubjectDetail(in *cmsclient.QuerySubjectDetailReq) (*cmsclient.QuerySubjectDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QuerySubjectDetailResp{}, nil
}
