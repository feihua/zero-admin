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
// 1.根据菜单名称查询菜单是否已存在
// 2.如果菜单已存在,则直接返回
// 3.菜单不存在时,则直接添加菜单
func (l *AddMenuLogic) AddMenu(in *sysclient.AddMenuReq) (*sysclient.AddMenuResp, error) {

	q := query.SysMenu.WithContext(l.ctx)
	// 1.根据菜单名称查询菜单是否已存在
	name := in.MenuName
	count, err := q.Where(query.SysMenu.MenuName.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据菜单名称：%s,查询菜单信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询菜单信息失败"))
	}

	//2.如果菜单已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "菜单信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("菜单：%s,已存在", name))
	}

	// 3.菜单不存在时,则直接添加菜单
	menu := &model.SysMenu{
		BackgroundURL: in.BackgroundUrl,
		CreateBy:      in.CreateBy,
		VueRedirect:   in.VueRedirect,
		VueIcon:       in.VueIcon,
		VuePath:       in.VuePath,
		MenuIcon:      in.MenuIcon,
		MenuName:      in.MenuName,
		MenuPath:      in.MenuPath,
		MenuPerms:     in.MenuPerms,
		MenuSort:      in.MenuSort,
		MenuStatus:    in.MenuStatus,
		MenuType:      in.MenuType,
		ParentID:      in.ParentId,
		Remark:        in.Remark,
		VueComponent:  in.VueComponent,
	}

	err = q.Create(menu)

	if err != nil {
		logc.Errorf(l.ctx, "新增菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("新增菜单信息失败")
	}

	return &sysclient.AddMenuResp{}, nil
}
