package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MenuUpdateLogic) MenuUpdate(req types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuUpdate(l.ctx, &sysclient.MenuUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Perms:    req.Perms,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		LastUpdateBy: "admin",
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
	}
	return &types.UpdateMenuResp{
		Code:    "000000",
		Message: "",
	}, nil
}
