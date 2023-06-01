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
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:         in.Id,
		Name:       in.Name,
		NickName:   sql.NullString{String: in.NickName},
		Avatar:     sql.NullString{String: in.Avatar},
		Password:   user.Password,
		Salt:       user.Salt,
		Email:      sql.NullString{String: in.Email},
		Mobile:     sql.NullString{String: in.Mobile},
		Status:     in.Status,
		DeptId:     in.DeptId,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy},
		UpdateTime: sql.NullTime{Time: time.Now()},
		DelFlag:    0,
		JobId:      in.JobId,
	})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.UserRoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserRoleModel.Insert(l.ctx, &sysmodel.SysUserRole{
		UserId:     in.Id,
		RoleId:     in.RoleId,
		CreateBy:   "admin",
		CreateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &sys.UserUpdateResp{}, nil
}
