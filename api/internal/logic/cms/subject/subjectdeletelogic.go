package subject

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectDeleteLogic {
	return &SubjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectDeleteLogic) SubjectDelete(req *types.DeleteSubjectReq) (resp *types.DeleteSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
