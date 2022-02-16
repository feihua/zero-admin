package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MenuAddLogic) MenuAdd(req types.AddMenuReq) (*types.AddMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuAdd(l.ctx, &sysclient.MenuAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Perms:    req.Perms,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		CreateBy:     "admin",
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功",
	}, nil
}
