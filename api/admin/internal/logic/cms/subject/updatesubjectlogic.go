package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectLogic {
	return &UpdateSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubjectLogic) UpdateSubject(req *types.UpdateSubjectReq) (resp *types.UpdateSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
