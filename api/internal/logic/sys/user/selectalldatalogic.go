package logic

import (
	"context"
	"strconv"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectAllDataLogic {
	return SelectAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllDataLogic) SelectAllData(req types.SelectDataReq) (*types.SelectDataResp, error) {
	roleList, roleErr := l.svcCtx.Sys.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if roleErr != nil {
		return nil, roleErr
	}

	var roleData []*types.RoleAllResp

	for _, role := range roleList.List {
		roleData = append(roleData, &types.RoleAllResp{
			Id:     role.Id,
			Name:   role.Name,
			Remark: role.Remark,
		})
	}

	jobList, jobErr := l.svcCtx.Sys.JobList(l.ctx, &sysclient.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if jobErr != nil {
		return nil, jobErr
	}

	var jobData []*types.JobAllResp

	for _, job := range jobList.List {
		jobData = append(jobData, &types.JobAllResp{
			Id:      job.Id,
			JobName: job.JobName,
		})
	}

	deptList, err := l.svcCtx.Sys.DeptList(l.ctx, &sysclient.DeptListReq{})

	if err != nil {
		return nil, errorx.NewDefaultError(" ")
	}

	var deptData []*types.DeptAllResp

	for _, dept := range deptList.List {
		deptData = append(deptData, &types.DeptAllResp{
			Id:       dept.Id,
			Value:    strconv.FormatInt(dept.Id, 10),
			Title:    dept.Name,
			ParentId: dept.ParentId,
		})
	}

	return &types.SelectDataResp{
		Code:    "000000",
		Message: "",
		RoleAll: roleData,
		DeptAll: deptData,
		JobAll:  jobData,
	}, nil
}
