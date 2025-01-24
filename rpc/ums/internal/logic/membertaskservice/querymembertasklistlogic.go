package membertaskservicelogic

import (
	"context"
	"errors"
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
Date: 2025/01/24 10:32:59
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
			Id:           item.ID,                                 //
			TaskName:     item.TaskName,                           // 任务名称
			TaskGrowth:   item.TaskGrowth,                         // 赠送成长值
			TaskIntegral: item.TaskIntegral,                       // 赠送积分
			TaskType:     item.TaskType,                           // 任务类型：0->新手任务；1->日常任务
			Status:       item.Status,                             // 状态：0->禁用；1->启用
			CreateBy:     item.CreateBy,                           // 创建者
			CreateTime:   time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:     item.UpdateBy,                           // 更新者
			UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	logc.Infof(l.ctx, "查询会员任务列表信息,参数：%+v,响应：%+v", in, list)

	return &umsclient.QueryMemberTaskListResp{
		Total: count,
		List:  list,
	}, nil
}
