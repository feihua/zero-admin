package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleListLogic 分配用户角色
/*
Author: LiuFeiHua
Date: 2024/5/23 17:38
*/
type UpdateUserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserRoleList 分配用户角色(角色id为1的是系统预留超级管理员角色,不用关联)
// 1.判断是否为超级管理员
// 2.删除用户与角色的关联
// 3.添加用户与角色的关联
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(in *sysclient.UpdateUserRoleListReq) (*sysclient.UpdateUserRoleListResp, error) {
	// 1.判断是否为超级管理员
	if in.UserId == 1 {
		return nil, errors.New("不允许操作超级管理员用户")
	}

	count, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.ID.Eq(in.UserId)).Count()
	if err != nil {
		return nil, errors.New("查询用户失败")
	}

	if count == 0 {
		return nil, errors.New("用户不存在")
	}

	err = query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysUserRole
		userId := in.UserId

		// 2.删除用户与角色的关联
		if _, err = q.WithContext(l.ctx).Where(q.UserID.Eq(userId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var userRoles []*model.SysUserRole
		for _, roleId := range in.RoleIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: roleId,
				UserID: userId,
			})
		}

		// 3.添加用户与角色的关联
		if err = q.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
			logc.Errorf(l.ctx, "添加用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("分配用户角色失败")
	}

	return &sysclient.UpdateUserRoleListResp{}, nil
}
