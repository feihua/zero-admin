package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *RoleUpdateLogic) RoleUpdate(in *us.RoleUpdateReq) (*us.RoleUpdateResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UsRolesModel.Update(usmodel.UsRoles{
		Id:       in.Data.Id,
		RoleName: sql.NullString{
			String: in.Data.RoleName,
			Valid:  true,
		},
		Remark:   sql.NullString{
			String: in.Data.Remark,
			Valid:  true,
		},
	})
	if err != nil{
		return &us.RoleUpdateResp{
			Result: false,
		}, err
	}

	return &us.RoleUpdateResp{
		Result: true,
	}, nil
}
