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

	count, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.ID.Eq(in.RoleId)).Count()

	if err != nil {
		return nil, errors.New("查询角色失败")
	}

	if count == 0 {
		return nil, errors.New("角色不存在")
	}

	if in.IsExist == 1 {
		return addRoleUser(in, l.ctx)
	} else if in.IsExist == 0 {
		return deleteRoleUser(in, l.ctx)
	} else {
		return nil, errors.New("参数错误,IsExist只能是0或者1")
	}

}

func deleteRoleUser(in *sysclient.CancelAuthorizationReq, ctx context.Context) (*sysclient.CancelAuthorizationResp, error) {
	userRole := query.SysUserRole
	q := userRole.WithContext(ctx)

	_, err := q.Where(userRole.RoleID.Eq(in.RoleId), userRole.UserID.In(in.UserIds...)).Delete()
	if err != nil {
		logc.Errorf(ctx, "取消授权失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("取消授权失败")
	}

	return &sysclient.CancelAuthorizationResp{}, nil
}

func addRoleUser(in *sysclient.CancelAuthorizationReq, ctx context.Context) (*sysclient.CancelAuthorizationResp, error) {
	q := query.SysUserRole.WithContext(ctx)
	for _, userId := range in.UserIds {
		count, err := q.Where(query.SysUserRole.UserID.Eq(userId), query.SysUserRole.RoleID.Eq(in.RoleId)).Count()
		if err != nil {
			return nil, errors.New("授权失败")
		}
		if count == 0 {
			err = q.Create(&model.SysUserRole{
				RoleID: in.RoleId,
				UserID: userId,
			})
			if err != nil {
				logc.Errorf(ctx, "授权失败,参数:%+v,异常:%s", in, err.Error())
				return nil, errors.New("授权失败")
			}
		}
	}

	return &sysclient.CancelAuthorizationResp{}, nil
}
