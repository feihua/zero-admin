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
// 2.菜单存在时,则直接更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *sysclient.UpdateMenuReq) (*sysclient.UpdateMenuResp, error) {
	q := query.SysMenu.WithContext(l.ctx)
	// 1.根据菜单id查询菜单是否已存在
	_, err := q.Where(query.SysMenu.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据菜单id：%d,查询菜单信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询菜单信息失败"))
	}

	// 2.菜单存在时,则直接更新菜单
	menu := &model.SysMenu{
		ID:            in.Id,
		BackgroundURL: in.BackgroundUrl,
		UpdateBy:      in.CreateBy,
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

	_, err = q.Updates(menu)

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", menu, err.Error())
		return nil, errors.New("更新菜单信息失败")
	}

	return &sysclient.UpdateMenuResp{}, nil
}
