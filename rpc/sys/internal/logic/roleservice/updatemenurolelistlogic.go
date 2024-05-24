package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuRoleListLogic 更新角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/24 15:31
*/
type UpdateMenuRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleListLogic {
	return &UpdateMenuRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenuRoleList 更新用户与角色的关联
func (l *UpdateMenuRoleListLogic) UpdateMenuRoleList(in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	//id为1的是系统预留超级管理员角色,不用关联
	if in.RoleId == 1 {
		return &sysclient.UpdateMenuRoleResp{}, nil
	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysRoleMenu
		if _, err := q.WithContext(l.ctx).Where(q.RoleID.Eq(in.RoleId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var roleMenus []*model.SysRoleMenu
		for _, menuId := range in.MenuIds {
			roleMenus = append(roleMenus, &model.SysRoleMenu{
				RoleID: in.RoleId,
				MenuID: menuId,
			})
		}

		if err := q.WithContext(l.ctx).CreateInBatches(roleMenus, len(roleMenus)); err != nil {
			logc.Errorf(l.ctx, "添加角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
