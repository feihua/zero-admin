package logic

import (
	"context"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *us.RoleDeleteReq) (*us.RoleDeleteResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UsRolesModel.Delete(in.Id)
	if err != nil{
		return &us.RoleDeleteResp{
			Result: false,
		}, err
	}
	// 删除 role id缓存
	DelRoleId(l.svcCtx.RedisConn, in.Id)

	return &us.RoleDeleteResp{
		Result: true,
	}, nil
}
