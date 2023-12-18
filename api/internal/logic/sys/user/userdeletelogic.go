package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 13:59
*/
type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserDelete 删除用户
func (l *UserDeleteLogic) UserDelete(req types.DeleteUserReq) (*types.DeleteUserResp, error) {

	_, err := l.svcCtx.UserService.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %d,删除用户异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除用户失败")
	}

	return &types.DeleteUserResp{
		Code:    "000000",
		Message: "删除用户成功",
	}, nil
}
