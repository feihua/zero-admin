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

// QueryMemberTaskDetailLogic 查询会员任务详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTaskDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskDetailLogic {
	return &QueryMemberTaskDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberTaskDetail 查询会员任务详情
func (l *QueryMemberTaskDetailLogic) QueryMemberTaskDetail(req *types.QueryMemberTaskDetailReq) (resp *types.QueryMemberTaskDetailResp, err error) {

	detail, err := l.svcCtx.MemberTaskService.QueryMemberTaskDetail(l.ctx, &umsclient.QueryMemberTaskDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberTaskDetailData{
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

	}
	return &types.QueryMemberTaskDetailResp{
		Code:    "000000",
		Message: "查询会员任务成功",
		Data:    data,
	}, nil
}
