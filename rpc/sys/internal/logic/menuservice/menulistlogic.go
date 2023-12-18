package menuservicelogic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

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
	count, _ := l.svcCtx.MenuModel.Count(l.ctx)
	all, err := l.svcCtx.MenuModel.FindAll(l.ctx, 1, count)

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*sysclient.MenuListData
	for _, menu := range *all {
		fmt.Println(menu)
		list = append(list, &sysclient.MenuListData{
			Id:             menu.Id,
			Name:           menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url.String,
			Perms:          menu.Perms.String,
			Type:           menu.Type,
			Icon:           menu.Icon.String,
			OrderNum:       menu.OrderNum.Int64,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   menu.UpdateBy.String,
			LastUpdateTime: menu.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        menu.DelFlag,
			VuePath:        menu.VuePath.String,
			VueComponent:   menu.VueComponent.String,
			VueIcon:        menu.VueIcon.String,
			VueRedirect:    menu.VueRedirect.String,
			BackgroundUrl:  menu.BackgroundUrl,
		})
	}

	logc.Infof(l.ctx, "查询菜单列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.MenuListResp{
		Total: count,
		List:  list,
	}, nil

}
