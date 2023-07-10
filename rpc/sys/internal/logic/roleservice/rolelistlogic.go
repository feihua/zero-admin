package roleservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

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

func (l *RoleListLogic) RoleList(in *sysclient.RoleListReq) (*sysclient.RoleListResp, error) {
	all, err := l.svcCtx.RoleModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.RoleModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询角色列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sysclient.RoleListData
	for _, role := range *all {
		list = append(list, &sysclient.RoleListData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark.String,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.UpdateBy.String,
			LastUpdateTime: role.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询角色列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sysclient.RoleListResp{
		Total: count,
		List:  list,
	}, nil

}
