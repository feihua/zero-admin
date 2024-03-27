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

// MenuUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:28
*/
type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuUpdateLogic {
	return MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuUpdate 更新菜单
func (l *MenuUpdateLogic) MenuUpdate(req types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	menuUpdateReq := sysclient.MenuUpdateReq{
		Id:            req.Id,
		Name:          req.Name,
		ParentId:      req.ParentId,
		Url:           req.Url,
		Perms:         req.Perms,
		Type:          req.Type,
		Icon:          req.Icon,
		OrderNum:      req.OrderNum,
		LastUpdateBy:  l.ctx.Value("userName").(string),
		VuePath:       req.VuePath,
		VueComponent:  req.VueComponent,
		VueIcon:       req.VueIcon,
		VueRedirect:   req.VueRedirect,
		DelFlag:       req.DelFlag,
		BackgroundUrl: req.BackgroundUrl,
	}
	if _, err := l.svcCtx.MenuService.MenuUpdate(l.ctx, &menuUpdateReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新菜单信息失败")
	}

	return &types.UpdateMenuResp{
		Code:    "000000",
		Message: "更新菜单信息成功!",
	}, nil
}
