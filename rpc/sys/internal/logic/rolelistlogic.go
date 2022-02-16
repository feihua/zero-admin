package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
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
	all, err := l.svcCtx.RoleModel.FindAll(in.Current, in.PageSize, in.Name)
	count, _ := l.svcCtx.RoleModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询角色列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sys.RoleListData
	for _, role := range *all {
		list = append(list, &sys.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询角色列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
