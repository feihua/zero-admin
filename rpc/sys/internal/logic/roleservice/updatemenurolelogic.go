package roleservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
)

type UpdateMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleLogic {
	return &UpdateMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	_ = l.svcCtx.RoleMenuModel.DeleteByRoleId(l.ctx, in.RoleId)

	for _, menuId := range in.MenuIds {
		_, _ = l.svcCtx.RoleMenuModel.Insert(l.ctx, &sysmodel.SysRoleMenu{
			RoleId:   in.RoleId,
			MenuId:   menuId,
			CreateBy: in.CreateBy,
		})
	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
