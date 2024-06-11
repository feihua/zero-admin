package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectStatusLogic {
	return &UpdateSubjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题表状态
func (l *UpdateSubjectStatusLogic) UpdateSubjectStatus(in *cmsclient.UpdateSubjectStatusReq) (*cmsclient.UpdateSubjectStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectStatusResp{}, nil
}
