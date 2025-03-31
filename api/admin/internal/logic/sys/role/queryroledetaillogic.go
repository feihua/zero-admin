package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleDetailLogic 查询角色详情
/*
Author: LiuFeiHua
Date: 2024/5/29 18:02
*/
type QueryRoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleDetail 查询角色详情
func (l *QueryRoleDetailLogic) QueryRoleDetail(req *types.QueryRoleDetailReq) (resp *types.QueryRoleDetailResp, err error) {
	detail, err := l.svcCtx.RoleService.QueryRoleDetail(l.ctx, &sysclient.QueryRoleDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询角色详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryRoleDetailData{
		Id:         detail.Id,         // 角色id
		RoleName:   detail.RoleName,   // 名称
		RoleKey:    detail.RoleKey,    // 角色权限字符串
		DataScope:  detail.DataScope,  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     detail.Status,     // 状态(1:正常，0:禁用)
		Remark:     detail.Remark,     // 备注
		DelFlag:    detail.DelFlag,    // 删除标志（0代表删除 1代表存在）
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间
	}

	return &types.QueryRoleDetailResp{
		Code:    "000000",
		Message: "查询角色详情成功",
		Data:    data,
	}, nil
}
