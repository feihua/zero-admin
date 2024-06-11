package tag

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberTagStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(req *types.UpdateMemberTagStatusReq) (resp *types.UpdateMemberTagStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
