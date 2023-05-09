package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {

	_ = l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:         in.Id,
		Name:       in.Name,
		NickName:   sql.NullString{String: in.NickName},
		Avatar:     sql.NullString{String: in.Avatar},
		Email:      sql.NullString{String: in.Email},
		Mobile:     sql.NullString{String: in.Mobile},
		DeptId:     in.DeptId,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy},
		Status:     in.Status,
		UpdateTime: sql.NullTime{Time: time.Now()},
		JobId:      in.JobId,
	})

	_ = l.svcCtx.UserRoleModel.Delete(l.ctx, in.Id)

	_, _ = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:     in.Id,
		RoleId:     in.RoleId,
		CreateBy:   "admin",
		CreateTime: time.Now(),
	})

	return &sys.UserUpdateResp{}, nil
}
