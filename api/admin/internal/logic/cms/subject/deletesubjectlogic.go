package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectLogic {
	return &DeleteSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSubjectLogic) DeleteSubject(req *types.DeleteSubjectReq) (resp *types.DeleteSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
