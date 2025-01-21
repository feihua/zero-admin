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

	if in.RoleId == 1 {
		return nil, errors.New("不允许操作超级管理员角色")
	}

	userRole := query.SysUserRole
	q := userRole.WithContext(l.ctx)

	// 确认授权
	if in.IsExist == 1 {
		var userRoles []*model.SysUserRole
		for _, userId := range in.UserIds {
			sysUserRole := model.SysUserRole{
				RoleID: in.RoleId,
				UserID: userId,
			}
			userRoles = append(userRoles, &sysUserRole)
		}

		// 2.添加角色与用户的关联
		err := q.CreateInBatches(userRoles, len(userRoles))

		if err != nil {
			logc.Errorf(l.ctx, "授权失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("授权失败")
		}
		return &sysclient.CancelAuthorizationResp{}, nil
	}

	// 取消授权
	_, err := q.Where(userRole.RoleID.Eq(in.RoleId), userRole.UserID.In(in.UserIds...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "取消授权失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("取消授权失败")
	}

	return &sysclient.CancelAuthorizationResp{}, nil
}
