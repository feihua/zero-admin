package subjectservicelogic

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectAddLogic {
	return &SubjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 专题
func (l *SubjectAddLogic) SubjectAdd(in *cmsclient.SubjectAddReq) (*cmsclient.SubjectAddResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.SubjectAddResp{}, nil
}
