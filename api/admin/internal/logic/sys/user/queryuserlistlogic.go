package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:04
*/
type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryUserListLogic {
	return QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserList 查询用户列表信息
func (l *QueryUserListLogic) QueryUserList(req *types.ListUserReq) (*types.ListUserResp, error) {

	resp, err := l.svcCtx.UserService.UserList(l.ctx, &sysclient.UserListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		UserName:   req.Name,
		NickName:   req.NickName,
		Mobile:     req.Mobile,
		Email:      req.Email,
		UserStatus: req.Status,
		DeptId:     req.DeptId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询用户列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询失败")
	}

	var list []*types.ListUserData

	for _, item := range resp.List {
		listUserData := types.ListUserData{}
		_ = copier.Copy(&listUserData, &item)
		list = append(list, &types.ListUserData{
			Id:         item.Id,
			UserName:   item.UserName,
			NickName:   item.NickName,
			Avatar:     item.Avatar,
			Email:      item.Email,
			Mobile:     item.Mobile,
			UserStatus: item.UserStatus,
			DeptId:     item.DeptId,
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime,
			UpdateBy:   item.UpdateBy,
			UpdateTime: item.UpdateTime,
			DelFlag:    item.DelFlag,
			JobId:      item.JobId,
			RoleId:     item.RoleId,
			RoleName:   item.RoleName,
			JobName:    item.JobName,
			DeptName:   item.DeptName,
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
