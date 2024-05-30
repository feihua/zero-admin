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

// AddMenuLogic 新增菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:24
*/
type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMenuLogic {
	return AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMenu 新增菜单
func (l *AddMenuLogic) AddMenu(req *types.AddMenuReq) (*types.AddMenuResp, error) {
	menuAddReq := sysclient.AddMenuReq{
		BackgroundUrl: req.BackgroundUrl,
		CreateBy:      l.ctx.Value("userName").(string),
		VueRedirect:   req.VueRedirect,
		VueIcon:       req.VueIcon,
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
		IsVisible:     req.IsVisible,
	}
	if _, err := l.svcCtx.MenuService.AddMenu(l.ctx, &menuAddReq); err != nil {
		logc.Errorf(l.ctx, "添加菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功!",
	}, nil
}
