package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteUserLogic 删除用户
/*
Author: LiuFeiHua
Date: 2023/12/18 14:20
*/
type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteUser 删除用户
// 1.判断是不是超级管理员
// 2.删除用户
// 3.删除用户与角色的关联
// 4.删除用户与岗位的关联
func (l *DeleteUserLogic) DeleteUser(in *sysclient.DeleteUserReq) (*sysclient.DeleteUserResp, error) {
	ids := in.Ids

	for _, roleId := range ids {
		// 1.判断是不是超级管理员
		if roleId == 1 {
			return nil, errors.New("删除用户失败,不允许操作超级管理员用户")
		}

	}

	err := query.Q.Transaction(func(tx *query.Query) error {
		// 2.删除用户
		user := tx.SysUser
		if _, err := user.WithContext(l.ctx).Where(user.ID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 3.删除用户与角色的关联
		role := tx.SysUserRole
		if _, err := role.WithContext(l.ctx).Where(role.UserID.In(ids...)).Delete(); err != nil {
			return err
		}

		// 4.删除用户与岗位的关联
		post := tx.SysUserPost
		if _, err := post.WithContext(l.ctx).Where(post.UserID.In(ids...)).Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除用户异常,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除用户异常")
	}

	return &sysclient.DeleteUserResp{}, nil
}
