package userservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:37
*/
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

// UserUpdate 更新用户(id为1是系统预留超级管理员用户,不能更新)
func (l *UserUpdateLogic) UserUpdate(in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {

	//id为1是系统预留超级管理员用户,不能更新
	if in.Id == 1 {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, "id为1是系统预留超级管理员用户,不能更新")
		return nil, errors.New("更新用户异常")
	}

	//查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, "用户不存在")
		return nil, errors.New("查询用户异常")
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
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("更新用户异常")
	}

	err = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, in.Id)
	if err != nil {
		logc.Errorf(l.ctx, "删除用户与角色关联异常,参数userId:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("更新用户异常")
	}

	_, err = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:   in.Id,
		RoleId:   in.RoleId,
		CreateBy: in.LastUpdateBy,
	})
	if err != nil {
		logc.Errorf(l.ctx, "添加用户与角色关联异常,参数userId:%d, roleId: %d,异常:%s", in.Id, in.RoleId, err.Error())
		return nil, errors.New("更新用户异常")
	}

	return &sysclient.UserUpdateResp{}, nil
}
