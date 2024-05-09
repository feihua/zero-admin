package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
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

	q := query.SysRoleMenu
	_, err := q.WithContext(l.ctx).Where(q.RoleID.Eq(in.RoleId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	for _, menuId := range in.MenuIds {
		_ = q.WithContext(l.ctx).Create(&model.SysRoleMenu{
			RoleID:   in.RoleId,
			MenuID:   menuId,
			CreateBy: in.CreateBy,
		})

	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
