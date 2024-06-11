package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectLogic {
	return &AddSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题表
func (l *AddSubjectLogic) AddSubject(in *cmsclient.AddSubjectReq) (*cmsclient.AddSubjectResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddSubjectResp{}, nil
}
