package roleservicelogic

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

// UpdateRoleLogic 更新角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:57
*/
type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRole 更新角色(id为1的是系统预留超级管理员角色,不能更新)
// 1.根据角色id查询角色是否已存在
// 2.角色存在时,则直接更新角色
func (l *UpdateRoleLogic) UpdateRole(in *sysclient.UpdateRoleReq) (*sysclient.UpdateRoleResp, error) {
	q := query.SysRole

	// 1.根据角色id查询角色是否已存在
	role, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色id：%d,查询角色信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询角色信息失败"))
	}

	if role.IsAdmin == 1 {
		logc.Errorf(l.ctx, "不能修改超级管理员的信息")
		return nil, errors.New(fmt.Sprintf("不能修改超级管理员角色"))
	}

	//3.角色存在时,则直接更新角色
	sysRole := &model.SysRole{
		ID:         in.Id,
		RoleName:   in.RoleName,
		RoleKey:    in.RoleKey,
		RoleStatus: in.RoleStatus,
		RoleSort:   in.RoleSort,
		DataScope:  in.DataScope,
		IsAdmin:    in.IsAdmin,
		Remark:     in.Remark,
		UpdateBy:   in.UpdateBy,
	}

	_, err = q.WithContext(l.ctx).Updates(sysRole)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", role, err.Error())
		return nil, errors.New("更新角色信息失败")
	}
	return &sysclient.UpdateRoleResp{}, nil
}
