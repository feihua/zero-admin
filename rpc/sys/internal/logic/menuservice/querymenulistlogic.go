package menuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuListLogic 菜单列表
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type QueryMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuListLogic {
	return &QueryMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMenuList 菜单列表
func (l *QueryMenuListLogic) QueryMenuList(in *sysclient.QueryMenuListReq) (*sysclient.QueryMenuListResp, error) {
	result, err := query.SysMenu.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询菜单列表信息失败")
	}
	var list []*sysclient.MenuListData
	for _, menu := range result {
		list = append(list, &sysclient.MenuListData{
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
		})
	}

	logc.Infof(l.ctx, "查询菜单列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryMenuListResp{
		Total: 0,
		List:  list,
	}, nil

}
