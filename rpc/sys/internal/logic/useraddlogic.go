package logic

import (
	"context"
	"go-zero-admin/rpc/model/sysmodel"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.UserAddResp, error) {

	_, _ = l.svcCtx.UserModel.Insert(sysmodel.SysUser{
		Name:           in.Name,
		NickName:       in.NickName,
		Avatar:         in.Avatar,
		Password:       "123456",
		Salt:           "123456",
		Email:          in.Email,
		Mobile:         in.Mobile,
		Status:         1,
		DeptId:         in.DeptId,
		CreateBy:       "admin",
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})

	return &sys.UserAddResp{}, nil
}
