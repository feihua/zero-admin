package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectAddLogic {
	return &SubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectAddLogic) SubjectAdd(req *types.AddSubjectReq) (resp *types.AddSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
