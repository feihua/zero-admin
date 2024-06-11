package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectDetailLogic {
	return &QuerySubjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySubjectDetailLogic) QuerySubjectDetail(req *types.QuerySubjectDetailReq) (resp *types.QuerySubjectDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
