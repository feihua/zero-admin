package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// RoleListLogic 查询角色列表
/*
Author: LiuFeiHua
Date: 2023/12/18 16:06
*/
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

// RoleList 查询角色列表
func (l *RoleListLogic) RoleList(in *sysclient.RoleListReq) (*sysclient.RoleListResp, error) {
	q := query.SysRole.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.SysRole.Name.Like("%" + in.Name + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.SysRole.Status.Eq(in.Status))
	}

	offset := (in.Current - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sysclient.RoleListData
	for _, role := range result {
		list = append(list, &sysclient.RoleListData{
			Id:         role.ID,
			Name:       role.Name,
			Remark:     role.Remark,
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:   role.UpdateBy,
			UpdateTime: common.TimeToString(role.UpdateTime),
			DelFlag:    role.DelFlag,
			Status:     role.Status,
		})
	}

	logc.Infof(l.ctx, "查询角色列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
