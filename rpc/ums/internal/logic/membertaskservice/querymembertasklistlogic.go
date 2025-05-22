package membertaskservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskListLogic 查询会员任务列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskListLogic {
	return &QueryMemberTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTaskList 查询会员任务列表
func (l *QueryMemberTaskListLogic) QueryMemberTaskList(in *umsclient.QueryMemberTaskListReq) (*umsclient.QueryMemberTaskListResp, error) {
	memberTask := query.UmsMemberTask
	q := memberTask.WithContext(l.ctx)
	if len(in.TaskName) > 0 {
		q = q.Where(memberTask.TaskName.Like("%" + in.TaskName + "%"))
	}

	if in.TaskType != 2 {
		q = q.Where(memberTask.TaskType.Eq(in.TaskType))
	}

	if in.RewardType != 2 {
		q = q.Where(memberTask.RewardType.Eq(in.RewardType))
	}

	// if in.StartTime != 2 {
	// 	q = q.Where(memberTask.StartTime.Eq(in.StartTime))
	// }
	// if in.EndTime != 2 {
	// 	q = q.Where(memberTask.EndTime.Eq(in.EndTime))
	// }
	if in.Status != 2 {
		q = q.Where(memberTask.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员任务列表失败")
	}

	var list []*umsclient.MemberTaskListData

	for _, item := range result {
		list = append(list, &umsclient.MemberTaskListData{
			Id:            item.ID,                                          // 主键ID
			TaskName:      item.TaskName,                                    // 任务名称
			TaskDesc:      item.TaskDesc,                                    // 任务描述
			TaskGrowth:    item.TaskGrowth,                                  // 赠送成长值
			TaskIntegral:  item.TaskIntegral,                                // 赠送积分
			TaskType:      item.TaskType,                                    // 任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务
			CompleteCount: item.CompleteCount,                               // 需要完成次数
			RewardType:    item.RewardType,                                  // 奖励类型：0-积分成长值，1-优惠券，2-抽奖次数
			RewardParams:  item.RewardParams,                                // 奖励参数JSON
			StartTime:     time_util.TimeToStr(item.StartTime),              // 任务开始时间
			EndTime:       time_util.TimeToStr(item.EndTime),                // 任务结束时间
			Status:        item.Status,                                      // 状态：0-禁用，1-启用
			Sort:          item.Sort,                                        // 排序
			CreateBy:      item.CreateBy,                                    // 创建人ID
			CreateTime:    time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:      pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:    time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &umsclient.QueryMemberTaskListResp{
		Total: count,
		List:  list,
	}, nil
}
