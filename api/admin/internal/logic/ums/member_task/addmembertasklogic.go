package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTaskLogic 添加会员任务
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskLogic {
	return &AddMemberTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberTask 添加会员任务
func (l *AddMemberTaskLogic) AddMemberTask(req *types.AddMemberTaskReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberTaskService.AddMemberTask(l.ctx, &umsclient.AddMemberTaskReq{
		TaskName:      req.TaskName,      // 任务名称
		TaskDesc:      req.TaskDesc,      // 任务描述
		TaskGrowth:    req.TaskGrowth,    // 赠送成长值
		TaskIntegral:  req.TaskIntegral,  // 赠送积分
		TaskType:      req.TaskType,      // 任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务
		CompleteCount: req.CompleteCount, // 需要完成次数
		RewardType:    req.RewardType,    // 奖励类型：0-积分成长值，1-优惠券，2-抽奖次数
		RewardParams:  req.RewardParams,  // 奖励参数JSON
		StartTime:     req.StartTime,     // 任务开始时间
		EndTime:       req.EndTime,       // 任务结束时间
		Status:        req.Status,        // 状态：0-禁用，1-启用
		Sort:          req.Sort,          // 排序
		CreateBy:      userId,            // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加会员任务成功",
	}, nil
}
