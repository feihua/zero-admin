package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/rpc/model/sysmodel"
	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"
	"time"
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
