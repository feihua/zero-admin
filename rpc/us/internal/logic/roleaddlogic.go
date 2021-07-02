package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/usmodel"
	"time"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *RoleAddLogic) RoleAdd(in *us.RoleAddReq) (*us.RoleAddResp, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UsRolesModel.Insert(usmodel.UsRoles{
		RoleName: sql.NullString{
			String: in.Data.RoleName,
			Valid:  true,
		},
		Remark: sql.NullString{
			String: in.Data.Remark,
			Valid:  true,
		},
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return &us.RoleAddResp{
			Result: false,
		}, err
	}
	// 增加 role缓存
	if id, err := result.LastInsertId(); err == nil {
		AddRoleId(l.svcCtx.RedisConn, id)
	}
	return &us.RoleAddResp{
		Result: true,
	}, nil
}
