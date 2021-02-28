package logic

import (
	"context"
	"go-zero-admin/rpc/model/sysmodel"
	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(in *sys.UpdateUserRoleReq) (*sys.UpdateUserRoleResp, error) {
	_ = l.svcCtx.UserRoleModel.Delete(in.Id)

	roleId, _ := strconv.ParseInt(in.RoleId, 10, 64)

	_, _ = l.svcCtx.UserRoleModel.Insert(sysmodel.SysUserRole{
		UserId:         in.Id,
		RoleId:         roleId,
		CreateBy:       "admin",
		CreateTime:     time.Now(),
		LastUpdateBy:   "admin",
		LastUpdateTime: time.Now(),
	})

	return &sys.UpdateUserRoleResp{}, nil
}
