package user

import (
	"context"
	"strconv"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryAllRelationsLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 13:39
*/
type QueryAllRelationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllRelationsLogic {
	return &QueryAllRelationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryAllRelations //查询用户的角色,单位,岗位的关系
func (l *QueryAllRelationsLogic) QueryAllRelations(req *types.QueryAllRelationsReq) (resp *types.QueryAllRelationsResp, err error) {
	//1.查询角色列表
	roleListResp, err := l.svcCtx.RoleService.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  1,
		PageSize: 100,
		Status:   1, //'状态  1:启用,0:禁用'
	})

	if err != nil {
		return nil, errorx.NewDefaultError("查询角色异常")
	}

	var roleRelations []*types.RoleRelations
	for _, role := range roleListResp.List {
		roleRelations = append(roleRelations, &types.RoleRelations{
			Id:     role.Id,
			Name:   role.Name,
			Remark: role.Remark,
		})
	}

	//2.查询岗位列表信息
	jobList, err := l.svcCtx.JobService.JobList(l.ctx, &sysclient.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		DelFlag:  1,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("查询岗位异常")
	}

	var jobRelations []*types.JobRelations
	for _, job := range jobList.List {
		jobRelations = append(jobRelations, &types.JobRelations{
			Id:      job.Id,
			JobName: job.JobName,
		})
	}

	//3.查询部门信息
	deptList, err := l.svcCtx.DeptService.DeptList(l.ctx, &sysclient.DeptListReq{})

	if err != nil {
		return nil, errorx.NewDefaultError("查询部门异常")
	}

	var deptRelations []*types.DeptRelations
	for _, dept := range deptList.List {
		deptRelations = append(deptRelations, &types.DeptRelations{
			Id:       dept.Id,
			Value:    strconv.FormatInt(dept.Id, 10),
			Title:    dept.Name,
			ParentId: dept.ParentId,
		})
	}

	return &types.QueryAllRelationsResp{
		Code:    "000000",
		Message: "查询成功",
		Data: types.QueryAllRelationsData{
			RoleRelations: roleRelations,
			DeptRelations: deptRelations,
			JobRelations:  jobRelations,
		},
	}, nil
}
