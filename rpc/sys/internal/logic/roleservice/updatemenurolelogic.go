package roleservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
)

// UpdateMenuRoleLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:58
*/
type UpdateMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleLogic {
	return &UpdateMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenuRole 更新角色与菜单的关联(id为1的是系统预留超级管理员角色,不用关联)
func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {

	//id为1的是系统预留超级管理员角色,不用关联
	if in.RoleId == 1 {
		return &sysclient.UpdateMenuRoleResp{}, nil
	}

	if err := l.svcCtx.RoleMenuModel.DeleteByRoleId(l.ctx, in.RoleId); err != nil {
		logc.Errorf(l.ctx, "删除角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	for _, menuId := range in.MenuIds {
		_, _ = l.svcCtx.RoleMenuModel.Insert(l.ctx, &sysmodel.SysRoleMenu{
			RoleId:   in.RoleId,
			MenuId:   menuId,
			CreateBy: in.CreateBy,
		})
	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
