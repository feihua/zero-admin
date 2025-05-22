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
	"gorm.io/gorm"
)

// QueryMemberTaskDetailLogic 查询会员任务详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTaskDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskDetailLogic {
	return &QueryMemberTaskDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTaskDetail 查询会员任务详情
func (l *QueryMemberTaskDetailLogic) QueryMemberTaskDetail(in *umsclient.QueryMemberTaskDetailReq) (*umsclient.QueryMemberTaskDetailResp, error) {
	item, err := query.UmsMemberTask.WithContext(l.ctx).Where(query.UmsMemberTask.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员任务不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员任务不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员任务异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员任务异常")
	}

	data := &umsclient.QueryMemberTaskDetailResp{
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
	}

	return data, nil
}
