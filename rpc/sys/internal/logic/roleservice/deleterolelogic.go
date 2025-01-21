package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteRoleLogic 删除角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:55
*/
type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteRole 删除角色(id为1的是系统预留超级管理员角色,不能删除)
// 1.判断是不是超级管理员
// 2.角色是否已使用
// 3.删除角色
// 4.删除用角色与菜单的关联
func (l *DeleteRoleLogic) DeleteRole(in *sysclient.DeleteRoleReq) (*sysclient.DeleteRoleResp, error) {
	ids := in.Ids

	for _, roleId := range ids {

		// 1.判断是不是超级管理员
		if roleId == 1 {
			return nil, errors.New("删除角色失败,不允许操作超级管理员角色")
		}

		// 2.角色是否已使用
		q := query.SysUserRole
		count, _ := q.WithContext(l.ctx).Select(q.RoleID).Where(q.RoleID.Eq(roleId)).Count()
		if count > 0 {
			return nil, errors.New("删除角色失败,已分配,不能删除")
		}

	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		// 3.删除角色
		role := tx.SysRole
		if _, err := role.WithContext(l.ctx).Where(role.ID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 4.删除用角色与菜单的关联
		menu := tx.SysRoleMenu
		if _, err := menu.WithContext(l.ctx).Where(menu.RoleID.In(ids...)).Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除角色信息失败")
	}

	return &sysclient.DeleteRoleResp{}, nil
}
