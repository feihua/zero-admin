package userservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:         in.Id,
		Name:       in.Name,
		NickName:   sql.NullString{String: in.NickName, Valid: true},
		Avatar:     sql.NullString{String: in.Avatar, Valid: true},
		Password:   user.Password,
		Salt:       user.Salt,
		Email:      sql.NullString{String: in.Email, Valid: true},
		Mobile:     sql.NullString{String: in.Mobile, Valid: true},
		Status:     in.Status,
		DeptId:     in.DeptId,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		JobId:      in.JobId,
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:   in.Id,
		RoleId:   in.RoleId,
		CreateBy: in.LastUpdateBy,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.UserUpdateResp{}, nil
}
