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

// UpdateMenuLogic 更新菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:28
*/
type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateMenuLogic {
	return UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMenu 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	menuReq := sysclient.UpdateMenuReq{
		BackgroundUrl: req.BackgroundUrl,
		CreateBy:      l.ctx.Value("userName").(string),
		VueRedirect:   req.VueRedirect,
		Id:            req.Id,
		VuePath:       req.VuePath,
		MenuIcon:      req.MenuIcon,
		MenuName:      req.MenuName,
		MenuPath:      req.MenuPath,
		MenuPerms:     req.MenuPerms,
		MenuSort:      req.MenuSort,
		MenuStatus:    req.MenuStatus,
		MenuType:      req.MenuType,
		ParentId:      req.ParentId,
		Remark:        req.Remark,
		VueComponent:  req.VueComponent,
		VueIcon:       req.VueIcon,
	}
	if _, err := l.svcCtx.MenuService.UpdateMenu(l.ctx, &menuReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateMenuResp{
		Code:    "000000",
		Message: "更新菜单信息成功!",
	}, nil
}
