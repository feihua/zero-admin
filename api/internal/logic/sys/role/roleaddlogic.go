package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleAddLogic {
	return RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req types.AddRoleReq) (*types.AddRoleResp, error) {
	_, err := l.svcCtx.Sys.RoleAdd(l.ctx, &sysclient.RoleAddReq{
		Name:   req.Name,
		Remark: req.Remark,
		//todo 从token里面拿
		CreateBy: "admin",
		Status:   req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加角色信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加角色失败")
	}

	return &types.AddRoleResp{
		Code:    "000000",
		Message: "添加角色成功",
	}, nil
}
