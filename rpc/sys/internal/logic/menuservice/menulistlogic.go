package menuservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuList 菜单列表
func (l *MenuListLogic) MenuList(in *sysclient.MenuListReq) (*sysclient.MenuListResp, error) {
	result, err := query.SysMenu.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*sysclient.MenuListData
	for _, menu := range result {
		list = append(list, &sysclient.MenuListData{
			Id:            menu.ID,
			Name:          menu.Name,
			ParentId:      menu.ParentID,
			Url:           menu.URL,
			Perms:         menu.Perms,
			Type:          menu.Type,
			Icon:          menu.Icon,
			OrderNum:      menu.OrderNum,
			CreateBy:      menu.CreateBy,
			CreateTime:    menu.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:      menu.UpdateBy,
			UpdateTime:    common.TimeToString(menu.UpdateTime),
			DelFlag:       menu.DelFlag,
			VuePath:       menu.VuePath,
			VueComponent:  menu.VueComponent,
			VueIcon:       menu.VueIcon,
			VueRedirect:   menu.VueRedirect,
			BackgroundUrl: menu.BackgroundURL,
		})
	}

	logc.Infof(l.ctx, "查询菜单列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.MenuListResp{
		Total: 0,
		List:  list,
	}, nil

}
