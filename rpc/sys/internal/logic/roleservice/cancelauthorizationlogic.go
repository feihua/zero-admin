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

// CancelAuthorizationLogic 取消授权/确认授权
/*
Author: LiuFeiHua
Date: 2024/6/02 18:31
*/
type CancelAuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAuthorizationLogic {
	return &CancelAuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelAuthorization 取消授权/确认授权
func (l *CancelAuthorizationLogic) CancelAuthorization(in *sysclient.CancelAuthorizationReq) (*sysclient.CancelAuthorizationResp, error) {

	//sys_role表is_admin为1表示系统预留超级管理员角色,不用关联
	role := query.SysRole
	count, _ := role.WithContext(l.ctx).Where(role.ID.Eq(in.RoleId), role.IsAdmin.Eq(1)).Count()
	if count == 1 {
		return &sysclient.CancelAuthorizationResp{}, nil
	}

	if in.IsExist == 1 {
		var userRoles []*model.SysUserRole
		for _, userId := range in.UserIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: in.RoleId,
				UserID: userId,
			})
		}

		// 2.添加角色与用户的关联
		err := query.SysUserRole.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles))

		if err != nil {
			logc.Errorf(l.ctx, "更新角色与用户的关联失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("更新角色与用户的关联失败")
		}
		return &sysclient.CancelAuthorizationResp{}, nil
	}

	userRole := query.SysUserRole
	_, err := userRole.WithContext(l.ctx).Where(userRole.RoleID.Eq(in.RoleId), userRole.UserID.In(in.UserIds...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "取消授权失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("取消授权失败")
	}

	return &sysclient.CancelAuthorizationResp{}, nil
}
