package menuservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuUpdateLogic 更新菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:46
*/
type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuUpdate 更新菜单
// 1.根据菜单id查询菜单是否已存在
// 2.如果菜单不存在,则直接返回
// 3.菜单存在时,则直接更新菜单
func (l *MenuUpdateLogic) MenuUpdate(in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	q := query.SysMenu.WithContext(l.ctx)
	// 1.根据菜单名称查询菜单是否已存在
	name := in.Name
	count, err := q.Where(query.SysMenu.Name.Eq(name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据菜单名称：%s,查询菜单信息失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询菜单信息失败"))
	}

	//2.如果菜单不存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "菜单信息不存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("菜单：%s,不存在", name))
	}

	// 3.菜单存在时,则直接更新菜单
	menu := &model.SysMenu{
		ID:            in.Id,
		Name:          in.Name,
		ParentID:      in.ParentId,
		URL:           in.Url,
		Perms:         in.Perms,
		Type:          in.Type,
		Icon:          in.Icon,
		OrderNum:      in.OrderNum,
		UpdateBy:      in.UpdateBy,
		DelFlag:       in.DelFlag,
		VuePath:       in.VuePath,
		VueComponent:  in.VueComponent,
		VueIcon:       in.VueIcon,
		VueRedirect:   in.VueRedirect,
		BackgroundURL: in.BackgroundUrl,
	}

	_, err = q.Updates(menu)

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("更新菜单信息失败")
	}

	return &sysclient.MenuUpdateResp{}, nil
}
