package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:24
*/
type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuAddLogic {
	return MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuAdd 新增菜单
func (l *MenuAddLogic) MenuAdd(req types.AddMenuReq) (*types.AddMenuResp, error) {
	menuAddReq := sysclient.MenuAddReq{
		Name:          req.Name,
		ParentId:      req.ParentId,
		Url:           req.Url,
		Perms:         req.Perms,
		Type:          req.Type,
		Icon:          req.Icon,
		OrderNum:      req.OrderNum,
		CreateBy:      l.ctx.Value("userName").(string),
		VuePath:       req.VuePath,
		VueComponent:  req.VueComponent,
		VueIcon:       req.VueIcon,
		VueRedirect:   req.VueRedirect,
		DelFlag:       req.DelFlag,
		BackgroundUrl: req.BackgroundUrl,
	}
	if _, err := l.svcCtx.MenuService.MenuAdd(l.ctx, &menuAddReq); err != nil {
		logc.Errorf(l.ctx, "添加菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功!",
	}, nil
}
