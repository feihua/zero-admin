package menuservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMenuLogic 新增菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:44
*/
type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMenu 新增菜单
// 1.查询菜单名称是否已存在,如果菜单已存在,则直接返回
// 2.查询菜单路由是否已存在,如果菜单已存在,则直接返回
// 3.菜单不存在时,则直接添加菜单
func (l *AddMenuLogic) AddMenu(in *sysclient.AddMenuReq) (*sysclient.AddMenuResp, error) {
	name := in.MenuName
	path := in.MenuPath

	q := query.SysMenu.WithContext(l.ctx)

	// 1.查询菜单名称是否已存在,如果菜单已存在,则直接返回
	count, err := q.Where(query.SysMenu.MenuName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单名称是否已存在失败, 参数：%s,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("新增菜单失败,菜单名称已存在"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("菜单：%s,已存在", name))
	}

	// 2.查询菜单路由是否已存在,如果菜单已存在,则直接返回
	count, err = q.Where(query.SysMenu.MenuPath.Eq(path)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单路由是否已存在失败,路由：%s,异常:%s", path, err.Error())
		return nil, errors.New(fmt.Sprintf("新增菜单失"))
	}

	if count > 0 {
		logc.Errorf(l.ctx, "路由已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("新增菜单失败,菜单路由：%s,已存在", path))
	}

	// 3.菜单不存在时,则直接添加菜单
	menu := &model.SysMenu{
		MenuName:      in.MenuName,      // 菜单名称
		ParentID:      in.ParentId,      // 父菜单ID，一级菜单为0
		MenuPath:      in.MenuPath,      // 前端路由
		MenuPerms:     in.MenuPerms,     // 权限标识
		MenuType:      in.MenuType,      // 类型 0：目录,1：菜单,2：按钮,3：外链
		MenuIcon:      in.MenuIcon,      // 菜单图标
		MenuSort:      in.MenuSort,      // 菜单排序
		CreateBy:      in.CreateBy,      // 创建者
		MenuStatus:    in.MenuStatus,    // 菜单状态
		IsDeleted:     in.IsDeleted,     // 是否删除  0：否  1：是
		IsVisible:     in.IsVisible,     // 是否可见  0：否  1：是
		Remark:        in.Remark,        // 备注信息
		VuePath:       in.VuePath,       // vue系统的path
		VueComponent:  in.VueComponent,  // vue的页面
		VueIcon:       in.VueIcon,       // vue的图标
		VueRedirect:   in.VueRedirect,   // vue的路由重定向
		BackgroundURL: in.BackgroundUrl, // 接口地址
	}

	err = q.Create(menu)

	if err != nil {
		logc.Errorf(l.ctx, "新增菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("新增菜单信息失败")
	}

	return &sysclient.AddMenuResp{}, nil
}
