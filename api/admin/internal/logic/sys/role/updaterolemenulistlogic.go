package role

import (
	"context"
	"strconv"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleMenuListLogic 更新角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:30
*/
type UpdateRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuListLogic {
	return &UpdateRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleMenuList 更新角色与菜单的关联
func (l *UpdateRoleMenuListLogic) UpdateRoleMenuList(req *types.UpdateRoleMenuListReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.RoleService.UpdateMenuRoleList(l.ctx, &sysclient.UpdateRoleMenuReq{
		RoleId:  req.RoleId,  // 角色id
		MenuIds: req.MenuIds, // 菜单id
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色菜单失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 以下操作是确保权限变更了,超级管理员不用重新登录
	key := "zero:mall:token"
	urls, err := l.svcCtx.Redis.HgetCtx(l.ctx, key, strconv.FormatInt(1, 10))
	if err != nil {
		logc.Error(l.ctx, "获取redis连接异常")
		return nil, errorx.NewDefaultError("获取redis连接异常")
	}
	fields, err := l.svcCtx.Redis.HkeysCtx(l.ctx, key)
	if err != nil {
		logc.Error(l.ctx, "获取redis连接异常")
		return nil, errorx.NewDefaultError("获取redis连接异常")
	}
	// 删除的时候 ,有可能修改了权限,所以需要清空除了管理员外,其它人权限
	_, err = l.svcCtx.Redis.HdelCtx(l.ctx, key, fields...)
	if err != nil {
		logc.Error(l.ctx, "获取redis连接异常")
		return nil, errorx.NewDefaultError("获取redis连接异常")
	}

	err = l.svcCtx.Redis.HsetCtx(l.ctx, key, strconv.FormatInt(1, 10), urls)
	if err != nil {
		logc.Error(l.ctx, "获取redis连接异常")
		return nil, errorx.NewDefaultError("获取redis连接异常")
	}
	return res.Success()
}
