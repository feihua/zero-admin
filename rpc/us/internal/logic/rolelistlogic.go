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

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *us.RoleListReq) (*us.RoleListResp, error) {
	// todo: add your logic here and delete this line
	all, _ := l.svcCtx.UsRolesModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UsRolesModel.Count()

	var list []*us.RoleData
	for _, item := range *all {
		list = append(list, &us.RoleData{
			Id:      item.Id,
			RoleName: item.RoleName.String,
			Remark:   item.Remark.String,
			CreateAt: item.CreateAt.Time.String(),
			UpdateAt: item.UpdateAt.Time.String(),
		})
	}

	return &us.RoleListResp{
		Total: count,
		List:  list,
	}, nil

	return &us.RoleListResp{}, nil
}
