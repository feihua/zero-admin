package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMenuLogic 新增菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:24
*/
type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMenuLogic {
	return AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMenu 新增菜单
func (l *AddMenuLogic) AddMenu(req *types.AddMenuReq) (*types.AddMenuResp, error) {
	createBy := l.ctx.Value("userName").(string)
	menuAddReq := sysclient.AddMenuReq{
		MenuName:      req.MenuName,      // 菜单名称
		ParentId:      req.ParentId,      // 父菜单ID，一级菜单为0
		MenuPath:      req.MenuPath,      // 前端路由
		MenuPerms:     req.MenuPerms,     // 权限标识
		MenuType:      req.MenuType,      // 类型 0：目录,1：菜单,2：按钮,3：外链
		MenuIcon:      req.MenuIcon,      // 菜单图标
		MenuSort:      req.MenuSort,      // 菜单排序
		CreateBy:      createBy,          // 创建者
		MenuStatus:    req.MenuStatus,    // 菜单状态
		IsVisible:     req.IsVisible,     // 是否可见  0：否  1：是
		Remark:        req.Remark,        // 备注信息
		VuePath:       req.VuePath,       // vue系统的path
		VueComponent:  req.VueComponent,  // vue的页面
		VueIcon:       req.VueIcon,       // vue的图标
		VueRedirect:   req.VueRedirect,   // vue的路由重定向
		BackgroundUrl: req.BackgroundUrl, // 接口地址
	}
	if _, err := l.svcCtx.MenuService.AddMenu(l.ctx, &menuAddReq); err != nil {
		logc.Errorf(l.ctx, "添加菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功!",
	}, nil
}
