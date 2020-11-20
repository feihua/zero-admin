package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, err := l.svcCtx.MenuModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
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
		})
	}

	return &sys.MenuListResp{
		Total: 10,
		List:  list,
	}, nil

}
