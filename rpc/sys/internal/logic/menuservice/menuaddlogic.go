package menuservicelogic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:44
*/
type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuAdd 新增菜单
func (l *MenuAddLogic) MenuAdd(in *sysclient.MenuAddReq) (*sysclient.MenuAddResp, error) {
	menu := &sysmodel.SysMenu{
		Name:          in.Name,
		ParentId:      in.ParentId,
		Url:           sql.NullString{String: in.Url, Valid: true},
		Perms:         sql.NullString{String: in.Perms, Valid: true},
		Type:          in.Type,
		Icon:          sql.NullString{String: in.Icon, Valid: true},
		OrderNum:      sql.NullInt64{Int64: in.OrderNum, Valid: true},
		CreateBy:      in.CreateBy,
		DelFlag:       in.DelFlag,
		VuePath:       sql.NullString{String: in.VuePath, Valid: true},
		VueComponent:  sql.NullString{String: in.VueComponent, Valid: true},
		VueIcon:       sql.NullString{String: in.VueIcon, Valid: true},
		VueRedirect:   sql.NullString{String: in.VueRedirect, Valid: true},
		BackgroundUrl: in.BackgroundUrl,
	}
	_, err := l.svcCtx.MenuModel.Insert(l.ctx, menu)

	if err != nil {
		logc.Errorf(l.ctx, "新增菜单信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.MenuAddResp{}, nil
}
