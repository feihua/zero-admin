package roleservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleUpdateLogic 更新角色
/*
Author: LiuFeiHua
Date: 2023/12/18 15:57
*/
type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleUpdate 更新角色(id为1的是系统预留超级管理员角色,不能更新)
// 1.根据角色id查询角色是否已存在
// 2.如果角色不存在,则直接返回
// 3.角色存在时,则直接更新角色
func (l *RoleUpdateLogic) RoleUpdate(in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {

	if in.Id == 1 {
		logc.Errorf(l.ctx, "不能修改超级管理员的信息")
		return nil, errors.New(fmt.Sprintf("不能修改超级管理员角色"))
	}

	q := query.SysRole.WithContext(l.ctx)
	// 1.根据角色名称查询角色是否已存在
	name := in.Name
	count, err := q.Where(query.SysRole.Name.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据角色名称：%s,查询角色信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询角色信息失败"))
	}

	//2.如果角色不存在,则直接返回
	if count == 0 {
		logc.Errorf(l.ctx, "角色信息不存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("角色：%s,不存在", name))
	}

	//3.角色存在时,则直接更新角色
	role := &model.SysRole{
		Name:     in.Name,
		Remark:   in.Remark,
		UpdateBy: in.UpdateBy,
		Status:   in.Status,
	}

	_, err = q.Updates(role)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", role, err.Error())
		return nil, errors.New("更新角色信息失败")
	}
	return &sysclient.RoleUpdateResp{}, nil
}
