package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleUserListLogic 根据roleId查询角色的用户
/*
Author: LiuFeiHua
Date: 2024/6/02 19:25
*/
type QueryRoleUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleUserListLogic {
	return &QueryRoleUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleUserList 根据roleId查询角色的用户
func (l *QueryRoleUserListLogic) QueryRoleUserList(req *types.QueryRoleUserListReq) (resp *types.QueryRoleUserListResp, err error) {
	result, err := l.svcCtx.RoleService.QueryRoleUserList(l.ctx, &sysclient.QueryRoleUserListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		RoleId:   req.RoleId,
		IsExist:  req.IsExist,
		Mobile:   req.Mobile,
		UserName: req.UserName,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryRoleUserListData

	for _, item := range result.List {
		list = append(list, &types.QueryRoleUserListData{
			Avatar:     item.Avatar,
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime,
			DeptId:     item.DeptId,
			Email:      item.Email,
			Id:         item.Id,
			LoginIp:    item.LoginIp,
			LoginTime:  item.LoginTime,
			Mobile:     item.Mobile,
			NickName:   item.NickName,
			Remark:     item.Remark,
			UpdateBy:   item.UpdateBy,
			UpdateTime: item.UpdateTime,
			UserName:   item.UserName,
			UserStatus: item.UserStatus,
		})
	}

	return &types.QueryRoleUserListResp{
		Code:     "000000",
		Message:  "查询角色用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil

}
