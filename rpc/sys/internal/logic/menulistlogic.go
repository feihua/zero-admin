package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MenuListLogic) MenuList(in *sys.MenuListReq) (*sys.MenuListResp, error) {
	count, _ := l.svcCtx.MenuModel.Count()
	all, err := l.svcCtx.MenuModel.FindAll(1, count)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询菜单列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*sys.MenuListData
	for _, menu := range *all {
		fmt.Println(menu)
		list = append(list, &sys.MenuListData{
			Id:             menu.Id,
			Name:           menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           menu.Type,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        menu.DelFlag,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询菜单列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.MenuListResp{
		Total: count,
		List:  list,
	}, nil

}
