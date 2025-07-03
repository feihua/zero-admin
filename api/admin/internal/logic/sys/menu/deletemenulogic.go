package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMenuLogic 删除菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:25
*/
type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMenuLogic {
	return DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMenu 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (*types.BaseResp, error) {
	if _, err := l.svcCtx.MenuService.DeleteMenu(l.ctx, &sysclient.DeleteMenuReq{
		Id: req.Id, // 编号
	}); err != nil {
		logc.Errorf(l.ctx, "根据menuId: %+v,删除菜单异常:%s", req, err.Error())
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

	// 删除的时候 ,有可能修改了权限,所以需要清空除了管理员外,其它人权限
	_, err = l.svcCtx.Redis.HdelCtx(l.ctx, key)
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
