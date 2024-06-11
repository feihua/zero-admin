package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectLogic {
	return &AddSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSubjectLogic) AddSubject(req *types.AddSubjectReq) (resp *types.AddSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
