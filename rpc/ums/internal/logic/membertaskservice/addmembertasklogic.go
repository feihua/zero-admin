package membertaskservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// AddMemberTaskLogic 添加会员任务
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type AddMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskLogic {
	return &AddMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTask 添加会员任务
func (l *AddMemberTaskLogic) AddMemberTask(in *umsclient.AddMemberTaskReq) (*umsclient.AddMemberTaskResp, error) {
	q := query.UmsMemberTask

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	item := &model.UmsMemberTask{
		TaskName:      in.TaskName,      // 任务名称
		TaskDesc:      in.TaskDesc,      // 任务描述
		TaskGrowth:    in.TaskGrowth,    // 赠送成长值
		TaskIntegral:  in.TaskIntegral,  // 赠送积分
		TaskType:      in.TaskType,      // 任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务
		CompleteCount: in.CompleteCount, // 需要完成次数
		RewardType:    in.RewardType,    // 奖励类型：0-积分成长值，1-优惠券，2-抽奖次数
		RewardParams:  in.RewardParams,  // 奖励参数JSON
		StartTime:     startTime,        // 任务开始时间
		EndTime:       endTime,          // 任务结束时间
		Status:        in.Status,        // 状态：0-禁用，1-启用
		Sort:          in.Sort,          // 排序
		CreateBy:      in.CreateBy,      // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员任务失败")
	}

	return &umsclient.AddMemberTaskResp{}, nil
}
