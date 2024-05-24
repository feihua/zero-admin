package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UserAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:13
*/
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

// UserAdd 新增用户
// 1.根据用户名称查询用户是否已存在
// 2.如果用户已存在,则直接返回
// 3.用户不存在时,则直接添加用户
func (l *UserAddLogic) UserAdd(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	user := &model.SysUser{
		UserName:   in.UserName,
		NickName:   in.NickName,
		Avatar:     in.Avatar,
		Password:   "123456",
		Salt:       "123456",
		Email:      in.Email,
		Mobile:     in.Mobile,
		UserStatus: in.UserStatus,
		DeptID:     in.DeptId,
		CreateBy:   in.CreateBy,
		DelFlag:    1,
		JobID:      in.JobId,
	}
	err := query.SysUser.WithContext(l.ctx).Create(user)

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("新增用户异常")
	}

	//获取新增用户的id
	id := user.ID

	//新增用户的时候,要删除它的关联信息(预防之前的数据导致查询权限的时候出问题)
	_, _ = query.SysUserRole.WithContext(l.ctx).Where(query.SysUserRole.UserID.Eq(id)).Delete()

	err = query.SysUserRole.WithContext(l.ctx).Create(&model.SysUserRole{
		UserID: id,
		//RoleID: in.RoleId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数userId:%d,roleId: %d,异常:%s", id, 1, err.Error())
		return nil, errors.New("添加用户与角色关联异常")
	}
	return &sysclient.UserAddResp{}, nil
}
