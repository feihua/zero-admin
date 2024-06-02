package roleservicelogic

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

// UpdateRoleUserListLogic 更新角色与用户的关联
/*
Author: LiuFeiHua
Date: 2024/6/02 18:31
*/
type UpdateRoleUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleUserListLogic {
	return &UpdateRoleUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRoleUserList 更新角色与用户的关联
// 1.删除角色与用户的关联
// 2.添加角色与用户的关联
func (l *UpdateRoleUserListLogic) UpdateRoleUserList(in *sysclient.UpdateRoleUserListReq) (*sysclient.UpdateRoleUserListResp, error) {
	//sys_role表is_admin为1表示系统预留超级管理员角色,不用关联
	role := query.SysRole
	count, _ := role.WithContext(l.ctx).Where(role.ID.Eq(in.RoleId), role.IsAdmin.Eq(1)).Count()
	if count == 1 {
		return &sysclient.UpdateRoleUserListResp{}, nil
	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysUserRole
		// 1.删除角色与用户的关联
		if _, err := q.WithContext(l.ctx).Where(q.RoleID.Eq(in.RoleId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除角色与用户的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var userRoles []*model.SysUserRole
		for _, userId := range in.UserIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: in.RoleId,
				UserID: userId,
			})
		}

		// 2.添加角色与用户的关联
		if err := q.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
			logc.Errorf(l.ctx, "添加角色与用户的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色与用户的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新角色与用户的关联失败")
	}

	return &sysclient.UpdateRoleUserListResp{}, nil
}
