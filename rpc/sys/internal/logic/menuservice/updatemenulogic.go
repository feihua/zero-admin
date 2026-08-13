package menuservicelogic

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuLogic 更新菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:46
*/
type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenu 更新菜单
// 1.根据菜单id查询菜单是否已存在
// 2.查询菜单名称是否已存在,如果菜单已存在,则直接返回
// 3.查询菜单路由是否已存在,如果菜单已存在,则直接返回
// 4.菜单存在时,则直接更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *sysclient.UpdateMenuReq) (*sysclient.UpdateMenuResp, error) {
	name := in.MenuName
	path := in.MenuUrl

	sMenu := query.SysMenu
	q := sMenu.WithContext(l.ctx)

	// 1.根据菜单id查询菜单是否已存在
	item, err := q.Where(query.SysMenu.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "菜单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("菜单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询菜单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("查询菜单异常")
	}

	// 2.查询菜单名称是否已存在,如果菜单已存在,则直接返回
	count, err := q.Where(sMenu.ID.Neq(in.Id), sMenu.MenuName.Eq(name)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询菜单名称是否已存在失败, 参数：%s,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("新增菜单失败,菜单名称已存在"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("菜单：%s,已存在", name))
	}

	// 3.查询菜单路由是否已存在,如果菜单已存在,则直接返回
	if len(path) != 0 {
		count, err = q.Where(sMenu.ID.Neq(in.Id), sMenu.MenuURL.Eq(path)).Count()

		if err != nil {
			logc.Errorf(l.ctx, "查询菜单路由是否已存在失败,路由：%s,异常:%s", path, err.Error())
			return nil, errors.New(fmt.Sprintf("新增菜单失"))
		}

		if count > 0 {
			logc.Errorf(l.ctx, "路由已存在：%+v", in)
			return nil, errors.New(fmt.Sprintf("新增菜单失败,菜单路由：%s,已存在", path))
		}
	}

	now := time.Now()
	menu := &model.SysMenu{
		ID:           in.Id,           // 主键
		MenuName:     in.MenuName,     // 菜单名称
		Ancestors:    in.Ancestors,    // 祖级列表
		MenuType:     in.MenuType,     // 菜单类型(1:目录,2:菜单,3:按钮)
		MenuURL:      in.MenuUrl,      // 路由路径
		MenuIcon:     in.MenuIcon,     // 菜单图标
		MenuSort:     in.MenuSort,     // 排序
		ParentID:     in.ParentId,     // 父id
		APIURL:       in.ApiUrl,       // 接口url
		Visible:      in.Visible,      // 显示状态（0:隐藏,显示:1）
		Status:       in.Status,       // 菜单状态(1:正常，0:禁用)
		Remark:       in.Remark,       // 备注
		VuePath:      in.VuePath,      // vue的path
		VueComponent: in.VueComponent, // vue的页面
		VueIcon:      in.VueIcon,      // vue的图标
		VueRedirect:  in.VueRedirect,  // vue的路由重定向
		AngularIcon:  in.AngularIcon,  // angular的图标
		ReactIcon:    in.ReactIcon,    // antd react的图标
		CreateBy:     item.CreateBy,   // 创建者
		CreateTime:   item.CreateTime, // 创建时间
		UpdateBy:     in.UpdateBy,     // 更新者
		UpdateTime:   &now,            // 更新时间
	}

	// 4.菜单存在时,则直接更新菜单
	err = l.svcCtx.DB.Model(&model.SysMenu{}).WithContext(l.ctx).Where(query.SysMenu.ID.Eq(in.Id)).Save(menu).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("更新菜单信息失败")
	}

	key := l.svcCtx.RedisKey + "menu"
	filed := strconv.FormatInt(in.Id, 10)
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, key, filed)
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, l.svcCtx.RedisKey+"api_url", in.ApiUrl)
	return &sysclient.UpdateMenuResp{}, nil
}
