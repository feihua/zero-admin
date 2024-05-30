package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuDetailLogic 查询菜单详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:54
*/
type QueryMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuDetail 查询菜单详情
func (l *QueryMenuDetailLogic) QueryMenuDetail(req *types.QueryMenuDetailReq) (resp *types.QueryMenuDetailResp, err error) {
	menu, err := l.svcCtx.MenuService.QueryMenuDetail(l.ctx, &sysclient.QueryMenuDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMenuDetailData{
		BackgroundUrl: menu.BackgroundUrl,
		CreateBy:      menu.CreateBy,
		CreateTime:    menu.CreateTime,
		Id:            menu.Id,
		MenuIcon:      menu.MenuIcon,
		MenuName:      menu.MenuName,
		MenuPath:      menu.MenuPath,
		MenuPerms:     menu.MenuPerms,
		MenuSort:      menu.MenuSort,
		MenuStatus:    menu.MenuStatus,
		MenuType:      menu.MenuType,
		ParentId:      menu.ParentId,
		Remark:        menu.Remark,
		UpdateBy:      menu.UpdateBy,
		UpdateTime:    menu.UpdateTime,
		VueComponent:  menu.VueComponent,
		VueIcon:       menu.VueIcon,
		VuePath:       menu.VuePath,
		VueRedirect:   menu.VueRedirect,
		IsVisible:     menu.IsVisible,
	}

	return &types.QueryMenuDetailResp{
		Code:    "000000",
		Message: "查询菜单详情成功",
		Data:    data,
	}, nil
}
