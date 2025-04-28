package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuLogic 更新菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:28
*/
type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateMenuLogic {
	return UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMenu 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (*types.BaseResp, error) {
	updateBy := l.ctx.Value("userName").(string)
	menuReq := sysclient.UpdateMenuReq{
		Id:            req.Id,
		MenuName:      req.MenuName,      // 菜单名称
		ParentId:      req.ParentId,      // 父菜单ID，一级菜单为0
		MenuPath:      req.MenuPath,      // 前端路由
		MenuPerms:     req.MenuPerms,     // 权限标识
		MenuType:      req.MenuType,      // 类型 0：目录,1：菜单,2：按钮,3：外链
		MenuIcon:      req.MenuIcon,      // 菜单图标
		MenuSort:      req.MenuSort,      // 菜单排序
		UpdateBy:      updateBy,          // 更新者
		MenuStatus:    req.MenuStatus,    // 菜单状态
		IsVisible:     req.IsVisible,     // 是否可见  0：否  1：是
		Remark:        req.Remark,        // 备注信息
		VuePath:       req.VuePath,       // vue系统的path
		VueComponent:  req.VueComponent,  // vue的页面
		VueIcon:       req.VueIcon,       // vue的图标
		VueRedirect:   req.VueRedirect,   // vue的路由重定向
		BackgroundUrl: req.BackgroundUrl, // 接口地址
	}
	if _, err := l.svcCtx.MenuService.UpdateMenu(l.ctx, &menuReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
