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
		Name:     req.Name,
		NickName: req.NickName,
		Mobile:   req.Mobile,
		Email:    req.Mobile,
		Status:   req.Status,
		DeptId:   req.DeptId,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListUserData

	for _, item := range resp.List {
		list = append(list, &types.ListUserData{
			Id:             item.Id,
			Name:           item.Name,
			NickName:       item.NickName,
			Password:       item.Password,
			Salt:           item.Salt,
			Email:          item.Email,
			Mobile:         item.Mobile,
			Status:         item.Status,
			DeptId:         item.DeptId,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime,
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime,
			DelFlag:        item.DelFlag,
		})
	}

	return &types.ListUserResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
