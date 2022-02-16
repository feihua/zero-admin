package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"
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

func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *sys.UpdateMenuRoleReq) (*sys.UpdateMenuRoleResp, error) {
	_ = l.svcCtx.RoleMenuModel.Delete(in.RoleId)

	ids := in.MenuIds
	for _, id := range ids {
		_, _ = l.svcCtx.RoleMenuModel.Insert(sysmodel.SysRoleMenu{
			RoleId:         in.RoleId,
			MenuId:         id,
			CreateBy:       "admin",
			CreateTime:     time.Now(),
			LastUpdateBy:   "admin",
			LastUpdateTime: time.Now(),
		})
	}

	return &sys.UpdateMenuRoleResp{}, nil
}
