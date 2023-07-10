package roleservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.RoleModel.Update(l.ctx, &sysmodel.SysRole{
		Id:         in.Id,
		Name:       in.Name,
		Remark:     sql.NullString{String: in.Remark, Valid: true},
		CreateBy:   role.CreateBy,
		CreateTime: role.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		DelFlag:    0,
		Status:     in.Status,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.RoleUpdateResp{}, nil
}
