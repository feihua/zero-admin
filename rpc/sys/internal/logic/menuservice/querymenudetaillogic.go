package menuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuDetailLogic 查询菜单详情
/*
Author: LiuFeiHua
Date: 2024/5/30 11:41
*/
type QueryMenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMenuDetail 查询菜单详情
func (l *QueryMenuDetailLogic) QueryMenuDetail(in *sysclient.QueryMenuDetailReq) (*sysclient.QueryMenuDetailResp, error) {
	menu, err := query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询菜单详情失败")
	}

	data := &sysclient.QueryMenuDetailResp{
		BackgroundUrl: menu.BackgroundURL,
		CreateBy:      menu.CreateBy,
		CreateTime:    menu.CreateTime.Format("2006-01-02 15:04:05"),
		Id:            menu.ID,
		VueRedirect:   menu.VueRedirect,
		MenuIcon:      menu.MenuIcon,
		MenuName:      menu.MenuName,
		MenuPath:      menu.MenuPath,
		MenuPerms:     menu.MenuPerms,
		MenuSort:      menu.MenuSort,
		MenuStatus:    menu.MenuStatus,
		MenuType:      menu.MenuType,
		ParentId:      menu.ParentID,
		Remark:        menu.Remark,
		UpdateBy:      menu.UpdateBy,
		UpdateTime:    time_util.TimeToString(menu.UpdateTime),
		VueComponent:  menu.VueComponent,
		VueIcon:       menu.VueIcon,
		VuePath:       menu.VuePath,
	}

	logc.Infof(l.ctx, "查询菜单详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
