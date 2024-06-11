package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectLogic {
	return &DeleteSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题表
func (l *DeleteSubjectLogic) DeleteSubject(in *cmsclient.DeleteSubjectReq) (*cmsclient.DeleteSubjectResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteSubjectResp{}, nil
}
