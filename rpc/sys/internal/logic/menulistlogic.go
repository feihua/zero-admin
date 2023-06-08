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
	count, _ := l.svcCtx.MenuModel.Count(l.ctx)
	all, err := l.svcCtx.MenuModel.FindAll(l.ctx, 1, count)

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询菜单列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.MenuListResp{
		Total: count,
		List:  list,
	}, nil

}
