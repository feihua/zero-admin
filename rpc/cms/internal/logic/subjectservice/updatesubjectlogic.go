package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectLogic {
	return &UpdateSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题表
func (l *UpdateSubjectLogic) UpdateSubject(in *cmsclient.UpdateSubjectReq) (*cmsclient.UpdateSubjectResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectResp{}, nil
}
