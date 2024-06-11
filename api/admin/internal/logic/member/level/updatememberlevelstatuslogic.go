package level

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLevelStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberLevelStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelStatusLogic {
	return &UpdateMemberLevelStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMemberLevelStatusLogic) UpdateMemberLevelStatus(req *types.UpdateMemberLevelStatusReq) (resp *types.UpdateMemberLevelStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
