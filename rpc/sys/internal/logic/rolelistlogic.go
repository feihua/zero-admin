package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

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

func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	all, _ := l.svcCtx.RoleModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.RoleModel.Count()

	var list []*sys.RoleListData
	for _, role := range *all {
		fmt.Println(role)
		list = append(list, &sys.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
		})
	}

	return &sys.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
