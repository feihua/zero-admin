package userservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserLogic 更新用户
/*
Author: LiuFeiHua
Date: 2023/12/18 14:37
*/
type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUser 更新用户(id为1是系统预留超级管理员用户,不能更新)
// 1.根据用户id查询用户是否已存在
// 2.用户存在时,则直接更新用户
func (l *UpdateUserLogic) UpdateUser(in *sysclient.UpdateUserReq) (*sysclient.UpdateUserResp, error) {

	role := query.SysUserRole
	count, _ := role.WithContext(l.ctx).Where(role.RoleID.Eq(1), role.UserID.Eq(in.Id)).Count()

	if count == 1 {
		logc.Errorf(l.ctx, "更新用户异常,参数userId:%d,异常:%s", in.Id, "id为1是系统预留超级管理员用户,不能更新")
		return nil, errors.New("不能修改超级管理员用户")
	}

	q := query.SysUser.WithContext(l.ctx)
	// 1.根据用户id查询用户是否已存在
	_, err := q.Where(query.SysUser.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户id：%d,查询用户信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询用户信息失败"))
	}

	// 2.用户存在时,则直接更新用户
	user := &model.SysUser{
		ID:         in.Id,
		UserName:   in.UserName,
		NickName:   in.NickName,
		Avatar:     in.Avatar,
		Email:      in.Email,
		Mobile:     in.Mobile,
		UserStatus: in.UserStatus,
		DeptID:     in.DeptId,
		Remark:     in.Remark,
		UpdateBy:   in.UpdateBy,
	}

	_, err = q.Updates(user)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户异常,参数:%+v,异常:%s", user, err.Error())
		return nil, errors.New("更新用户失败")
	}

	return &sysclient.UpdateUserResp{}, nil
}
