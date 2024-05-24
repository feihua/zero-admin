package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

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
// 1.根据用户id查询用户是否已存在
// 2.如果用户不已存在,则直接返回
// 3.用户存在时,则直接更新用户
func (l *UserUpdateLogic) UserUpdate(in *sysclient.UserUpdateReq) (*sysclient.UserUpdateResp, error) {

	//id为1是系统预留超级管理员用户,不能更新
	if in.Id == 1 {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, "id为1是系统预留超级管理员用户,不能更新")
		return nil, errors.New("更新用户异常")
	}

	//查询用户是否存在
	q := query.SysUser
	user, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, "用户不存在")
		return nil, errors.New("查询用户异常")
	}

	sysUser := &model.SysUser{
		ID:         in.Id,
		UserName:   in.UserName,
		NickName:   in.NickName,
		Avatar:     in.Avatar,
		Password:   user.Password,
		Salt:       user.Salt,
		Email:      in.Email,
		Mobile:     in.Mobile,
		UserStatus: in.UserStatus,
		DeptID:     in.DeptId,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime,
		UpdateBy:   in.UpdateBy,
		JobID:      in.JobId,
	}

	_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Updates(sysUser)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("更新用户异常")
	}

	_, err = query.SysUserRole.WithContext(l.ctx).Where(query.SysUserRole.UserID.Eq(in.Id)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除用户与角色关联异常,参数userId:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("更新用户异常")
	}

	err = query.SysUserRole.WithContext(l.ctx).Create(&model.SysUserRole{
		UserID: in.Id,
		//RoleID: in.RoleId,
	})
	if err != nil {
		logc.Errorf(l.ctx, "添加用户与角色关联异常,参数userId:%d, roleId: %d,异常:%s", in.Id, 1, err.Error())
		return nil, errors.New("更新用户异常")
	}

	return &sysclient.UserUpdateResp{}, nil
}
