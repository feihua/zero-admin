package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleUpdateLogic
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
func (l *RoleUpdateLogic) RoleUpdate(in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {
	q := query.SysRole
	//id为1的是系统预留超级管理员角色,不用关联
	if in.Id == 1 {
		return &sysclient.RoleUpdateResp{}, nil
	}

	role, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	_, err = q.WithContext(l.ctx).Updates(&model.SysRole{
		ID:         in.Id,
		Name:       in.Name,
		Remark:     &in.Remark,
		CreateBy:   role.CreateBy,
		CreateTime: role.CreateTime,
		UpdateBy:   &in.UpdateBy,
		DelFlag:    role.DelFlag,
		Status:     in.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.RoleUpdateResp{}, nil
}
