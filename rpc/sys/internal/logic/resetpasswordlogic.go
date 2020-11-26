package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(in *sys.ReSetPasswordReq) (*sys.ReSetPasswordResp, error) {

	_ = l.svcCtx.UserModel.Update(model.SysUser{
		Id:             in.Id,
		Password:       "123456",
		Salt:           "123456",
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
	})

	return &sys.ReSetPasswordResp{}, nil
}
