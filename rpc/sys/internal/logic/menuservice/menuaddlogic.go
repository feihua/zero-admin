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

// MenuAddLogic 新增菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:44
*/
type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuAdd 新增菜单
// 1.根据菜单名称查询菜单是否已存在
// 2.如果菜单已存在,则直接返回
// 3.菜单不存在时,则直接添加菜单
func (l *MenuAddLogic) MenuAdd(in *sysclient.MenuAddReq) (*sysclient.MenuAddResp, error) {

	q := query.SysMenu.WithContext(l.ctx)
	// 1.根据菜单名称查询菜单是否已存在
	name := in.Name
	count, err := q.Where(query.SysMenu.Name.Eq(name)).Count()

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
		Name:          in.Name,
		ParentID:      in.ParentId,
		URL:           in.Url,
		Perms:         in.Perms,
		Type:          in.Type,
		Icon:          in.Icon,
		OrderNum:      in.OrderNum,
		CreateBy:      in.CreateBy,
		DelFlag:       in.DelFlag,
		VuePath:       in.VuePath,
		VueComponent:  in.VueComponent,
		VueIcon:       in.VueIcon,
		VueRedirect:   in.VueRedirect,
		BackgroundURL: in.BackgroundUrl,
	}

	err = q.Create(menu)

	if err != nil {
		logc.Errorf(l.ctx, "新增菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, err
	}

	return &sysclient.MenuAddResp{}, nil
}
