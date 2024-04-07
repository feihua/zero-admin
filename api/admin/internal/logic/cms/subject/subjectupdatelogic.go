package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectUpdateLogic {
	return &SubjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectUpdateLogic) SubjectUpdate(req *types.UpdateSubjectReq) (resp *types.UpdateSubjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
