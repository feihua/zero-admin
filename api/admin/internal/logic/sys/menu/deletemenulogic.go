package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMenuLogic 删除菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:25
*/
type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMenuLogic {
	return DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMenu 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (*types.BaseResp, error) {
	if _, err := l.svcCtx.MenuService.DeleteMenu(l.ctx, &sysclient.DeleteMenuReq{
		Id: req.Id, // 编号
	}); err != nil {
		logc.Errorf(l.ctx, "根据menuId: %+v,删除菜单异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
