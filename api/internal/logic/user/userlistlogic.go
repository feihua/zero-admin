package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserListLogic {
	return UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req types.ListUserReq) (*types.ListUserResp, error) {
	resp, err := l.svcCtx.Sys.UserList(l.ctx, &sysclient.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListUserData

	for _, user := range resp.List {
		list = append(list, &types.ListUserData{
			Id:             user.Id,
			Name:           user.Name,
			NickName:       user.NickName,
			Password:       user.Password,
			Salt:           user.Salt,
			Email:          user.Email,
			Mobile:         user.Mobile,
			DeptId:         user.DeptId,
			CreateBy:       user.CreateBy,
			CreateTime:     user.CreateTime,
			LastUpdateBy:   user.LastUpdateBy,
			LastUpdateTime: user.LastUpdateTime,
			DelFlag:        user.DelFlag,
		})
	}

	return &types.ListUserResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
