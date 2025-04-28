package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteUserLogic 删除用户
/*
Author: LiuFeiHua
Date: 2023/12/18 13:59
*/
type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteUser 删除用户
func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (*types.BaseResp, error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	for _, id := range req.Ids {
		if id == userId {
			return nil, errors.New("删除用户信息失败,当前用户不能删除")
		}

	}
	_, err := l.svcCtx.UserService.DeleteUser(l.ctx, &sysclient.DeleteUserReq{
		Ids: req.Ids, // 用户id
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %+v,删除用户异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
