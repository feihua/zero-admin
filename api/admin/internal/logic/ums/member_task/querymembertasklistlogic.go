package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskListLogic 查询会员任务列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
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
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		TaskName:   req.TaskName,   // 任务名称
		TaskType:   req.TaskType,   // 任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务
		RewardType: req.RewardType, // 奖励类型：0-积分成长值，1-优惠券，2-抽奖次数
		StartTime:  req.StartTime,  // 任务开始时间
		EndTime:    req.EndTime,    // 任务结束时间
		Status:     req.Status,     // 状态：0-禁用，1-启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字会员任务列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberTaskListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberTaskListData{
			Id:            detail.Id,            // 主键ID
			TaskName:      detail.TaskName,      // 任务名称
			TaskDesc:      detail.TaskDesc,      // 任务描述
			TaskGrowth:    detail.TaskGrowth,    // 赠送成长值
			TaskIntegral:  detail.TaskIntegral,  // 赠送积分
			TaskType:      detail.TaskType,      // 任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务
			CompleteCount: detail.CompleteCount, // 需要完成次数
			RewardType:    detail.RewardType,    // 奖励类型：0-积分成长值，1-优惠券，2-抽奖次数
			RewardParams:  detail.RewardParams,  // 奖励参数JSON
			StartTime:     detail.StartTime,     // 任务开始时间
			EndTime:       detail.EndTime,       // 任务结束时间
			Status:        detail.Status,        // 状态：0-禁用，1-启用
			Sort:          detail.Sort,          // 排序
			CreateBy:      detail.CreateBy,      // 创建人ID
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateBy:      detail.UpdateBy,      // 更新人ID
			UpdateTime:    detail.UpdateTime,    // 更新时间

		})
	}

	return &types.QueryMemberTaskListResp{
		Code:     "000000",
		Message:  "查询会员任务列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
