package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleListLogic {
	return RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req types.ListRoleReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.Sys.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		DelFlag:  req.DelFlag,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListRoleData

	for _, role := range resp.List {
		list = append(list, &types.ListRoleData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime,
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime,
			DelFlag:        role.DelFlag,
			Label:          role.Name,
			Value:          strconv.FormatInt(role.Id, 10),
		})
	}

	return &types.ListRoleResp{
		Code:     "000000",
		Message:  "",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
