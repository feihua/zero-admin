package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllRoleLogic {
	return AllRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllRoleLogic) AllRole(req types.AllRoleReq) (*types.AllRoleResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Us.AllRoles(l.ctx, &usclient.AllRoleReq{
	})
	if err != nil {
		return nil, err
	}

	allRoleData := types.AllRoleData{
		Total: resp.Total,
	}
	for _, v := range resp.List {
		roleData := types.RoleData{
			Id:       v.Id,
			RoleName: v.RoleName,
			Remark:   v.Remark,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		allRoleData.RoleData = append(allRoleData.RoleData, &roleData)
	}

	return &types.AllRoleResp{
		Code:    0,
		Message: "Success",
		Data:    allRoleData,
	}, nil
}
