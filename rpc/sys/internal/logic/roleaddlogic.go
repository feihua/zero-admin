package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	_, _ = l.svcCtx.RoleModel.Insert(sysmodel.SysRole{
		Id:             0,
		Name:           in.Name,
		Remark:         in.Remark,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Time{},
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		Status:         in.Status,
	})

	return &sys.RoleAddResp{}, nil
}
