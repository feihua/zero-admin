package task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskListLogic 查询会员任务列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:14
*/
type QueryMemberTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskListLogic {
	return &QueryMemberTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberTaskList 查询会员任务列表
func (l *QueryMemberTaskListLogic) QueryMemberTaskList(req *types.QueryMemberTaskListReq) (resp *types.QueryMemberTaskListResp, err error) {
	result, err := l.svcCtx.MemberTaskService.QueryMemberTaskList(l.ctx, &umsclient.QueryMemberTaskListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员任务列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员任务失败")
	}

	var list []*types.QueryMemberTaskListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberTaskListData{
			Id:           item.Id,
			TaskName:     item.TaskName,
			TaskGrowth:   item.TaskGrowth,
			TaskIntegral: item.TaskIntegral,
			TaskType:     item.TaskType,
		})
	}

	return &types.QueryMemberTaskListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员任务列表成功",
	}, nil
}
