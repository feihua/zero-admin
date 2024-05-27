package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleListLogic 查询角色列表
/*
Author: LiuFeiHua
Date: 2023/12/18 15:39
*/
type QueryRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryRoleListLogic {
	return QueryRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleList 查询角色列表
func (l *QueryRoleListLogic) QueryRoleList(req *types.ListRoleReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.RoleService.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListRoleData

	for _, role := range resp.List {
		list = append(list, &types.ListRoleData{
			Id:         role.Id,
			Name:       role.Name,
			Remark:     role.Remark,
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime,
			UpdateBy:   role.UpdateBy,
			UpdateTime: role.UpdateTime,
			DelFlag:    role.DelFlag,
			Label:      role.Name,
			Value:      strconv.FormatInt(role.Id, 10),
			Status:     role.Status,
		})
	}

	return &types.ListRoleResp{
		Code:     "000000",
		Message:  "查询角色成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
