package logic

import (
	"context"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var cacheUsRolesPrefix = "cache#usRoles#all#"

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *us.RoleListReq) (*us.RoleListResp, error) {

	var list []*us.RoleData

	all, _ := l.svcCtx.UsRolesModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UsRolesModel.Count()
	for _, item := range *all {
		list = append(list, &us.RoleData{
			Id:       item.Id,
			RoleName: item.RoleName.String,
			Remark:   item.Remark.String,
			CreateTime: item.CreateTime.Time.String(),
			UpdateTime: item.UpdateTime.Time.String(),
		})
	}

	return &us.RoleListResp{
		Total: count,
		List:  list,
	}, nil
}




