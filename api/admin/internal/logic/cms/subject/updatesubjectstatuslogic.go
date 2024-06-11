package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectStatusLogic {
	return &UpdateSubjectStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubjectStatusLogic) UpdateSubjectStatus(req *types.UpdateSubjectStatusReq) (resp *types.UpdateSubjectStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
