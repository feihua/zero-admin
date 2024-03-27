package configservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigRoleLogic {
	return &UpdateConfigRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigRoleLogic) UpdateConfigRole(in *sysclient.UpdateConfigRoleReq) (*sysclient.UpdateConfigRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateConfigRoleResp{}, nil
}
