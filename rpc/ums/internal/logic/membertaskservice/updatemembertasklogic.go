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
	"gorm.io/gorm"
	"time"
)

// UpdateMemberTaskLogic 更新会员任务
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type UpdateMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTaskLogic {
	return &UpdateMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTask 更新会员任务
func (l *UpdateMemberTaskLogic) UpdateMemberTask(in *umsclient.UpdateMemberTaskReq) (*umsclient.UpdateMemberTaskResp, error) {
	q := query.UmsMemberTask.WithContext(l.ctx)

	// 1.根据会员任务id查询会员任务是否已存在
	task, err := q.Where(query.UmsMemberTask.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员任务不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员任务不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员任务异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员任务异常")
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	now := time.Now()
	item := &model.UmsMemberTask{
		ID:            in.Id,            // 主键ID
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
		CreateBy:      task.CreateBy,    // 创建人ID
		CreateTime:    task.CreateTime,  // 创建时间
		UpdateBy:      &in.UpdateBy,     // 更新人ID
		UpdateTime:    &now,             // 更新时间
	}

	// 2.会员任务存在时,则直接更新会员任务
	err = l.svcCtx.DB.Model(&model.UmsMemberTask{}).WithContext(l.ctx).Where(query.UmsMemberTask.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新会员任务失败")
	}

	return &umsclient.UpdateMemberTaskResp{}, nil
}
