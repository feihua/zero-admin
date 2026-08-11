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
	path := in.MenuUrl

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
	if len(path) != 0 {
		count, err = q.Where(query.SysMenu.MenuURL.Eq(path)).Count()

		if err != nil {
			logc.Errorf(l.ctx, "查询菜单路由是否已存在失败,路由：%s,异常:%s", path, err.Error())
			return nil, errors.New(fmt.Sprintf("新增菜单失"))
		}

		if count > 0 {
			logc.Errorf(l.ctx, "路由已存在：%+v", in)
			return nil, errors.New(fmt.Sprintf("新增菜单失败,菜单路由：%s,已存在", path))
		}
	}

	// 3.菜单不存在时,则直接添加菜单
	menu := &model.SysMenu{
		MenuName: in.MenuName, // 菜单名称
		// Ancestors:    in.Ancestors,    // 祖级列表
		MenuType:     in.MenuType,     // 菜单类型(1:目录,2:菜单,3:按钮)
		MenuURL:      in.MenuUrl,      // 路由路径
		MenuIcon:     in.MenuIcon,     // 菜单图标
		MenuSort:     in.MenuSort,     // 排序
		ParentID:     in.ParentId,     // 父id
		APIURL:       in.ApiUrl,       // 接口url
		Visible:      in.Visible,      // 显示状态（0:隐藏,显示:1）
		Status:       in.Status,       // 菜单状态(1:正常，0:禁用)
		DelFlag:      1,               // 删除标志（0:删除,1:存在）
		Remark:       in.Remark,       // 备注
		VuePath:      in.VuePath,      // vue的path
		VueComponent: in.VueComponent, // vue的页面
		VueIcon:      in.VueIcon,      // vue的图标
		VueRedirect:  in.VueRedirect,  // vue的路由重定向
		AngularIcon:  in.AngularIcon,  // angular的图标
		ReactIcon:    in.ReactIcon,    // antd react的图标
		CreateBy:     in.CreateBy,     // 创建者
	}

	err = q.Create(menu)

	if err != nil {
		logc.Errorf(l.ctx, "新增菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("新增菜单信息失败")
	}

	return &sysclient.AddMenuResp{}, nil
}
