package menuservicelogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuDelete 删除菜单
func (l *MenuDeleteLogic) MenuDelete(in *sysclient.MenuDeleteReq) (*sysclient.MenuDeleteResp, error) {
	err := l.svcCtx.MenuModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		logc.Errorf(l.ctx, "删除菜单信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.MenuDeleteResp{}, nil
}
