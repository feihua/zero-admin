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

// UpdateMenuStatusLogic 更新菜单状态
/*
Author: LiuFeiHua
Date: 2024/5/29 17:50
*/
type UpdateMenuStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuStatusLogic {
	return &UpdateMenuStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMenuStatus 更新菜单状态
func (l *UpdateMenuStatusLogic) UpdateMenuStatus(req *types.UpdateMenuStatusReq) (resp *types.UpdateMenuStatusResp, err error) {
	menuReq := sysclient.UpdateMenuStatusReq{
		Ids:        req.MenuIds,
		MenuStatus: req.MenuStatus,
		UpdateBy:   l.ctx.Value("userName").(string),
	}
	if _, err = l.svcCtx.MenuService.UpdateMenuStatus(l.ctx, &menuReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单状态失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateMenuStatusResp{
		Code:    "000000",
		Message: "更新菜单状态成功!",
	}, nil
}
