package userservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {

	insert, _ := l.svcCtx.UserModel.Insert(l.ctx, &sysmodel.SysUser{
		Name:     in.Name,
		NickName: sql.NullString{String: in.NickName, Valid: true},
		Avatar:   sql.NullString{String: in.Avatar, Valid: true},
		Password: "123456",
		Salt:     "123456",
		Email:    sql.NullString{String: in.Email, Valid: true},
		Mobile:   sql.NullString{String: in.Mobile, Valid: true},
		Status:   in.Status,
		DeptId:   in.DeptId,
		CreateBy: in.CreateBy,
		DelFlag:  0,
		JobId:    in.JobId,
	})

	id, _ := insert.LastInsertId()
	_ = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, id)

	_, _ = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:   id,
		RoleId:   in.RoleId,
		CreateBy: in.CreateBy,
	})

	return &sysclient.UserAddResp{}, nil
}
