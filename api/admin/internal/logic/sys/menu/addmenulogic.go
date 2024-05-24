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

// AddMenuLogic
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
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功!",
	}, nil
}
