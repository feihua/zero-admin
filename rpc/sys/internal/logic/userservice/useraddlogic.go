package userservicelogic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/feihua/zero-admin/rpc/model/sysmodel"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

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
func (l *UserAddLogic) UserAdd(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {

	user := &sysmodel.SysUser{
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
	}
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, user)

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("新增用户异常")
	}

	//获取新增用户的id
	id, _ := insert.LastInsertId()

	//新增用户的时候,要删除它的关联信息(预防之前的数据导致查询权限的时候出问题)
	_ = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, id)

	_, err = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:   id,
		RoleId:   in.RoleId,
		CreateBy: in.CreateBy,
	})

	if err != nil {
		logc.Errorf(l.ctx, "新增用户异常,参数userId:%d,roleId: %d,异常:%s", id, in.RoleId, err.Error())
		return nil, errors.New("添加用户与角色关联异常")
	}
	return &sysclient.UserAddResp{}, nil
}
