package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/logic/common"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskListLogic 查询会员任务列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:52
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

// QueryMemberTaskList 查询会员任务表列表
func (l *QueryMemberTaskListLogic) QueryMemberTaskList(in *umsclient.QueryMemberTaskListReq) (*umsclient.QueryMemberTaskListResp, error) {
	q := query.UmsMemberTask.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberTaskListData
	for _, item := range result {

		list = append(list, &umsclient.MemberTaskListData{
			Id:           item.ID,
			TaskName:     item.TaskName,
			TaskGrowth:   item.TaskGrowth,
			TaskIntegral: item.TaskIntegral,
			TaskType:     item.TaskType,
			CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"),
			CreateBy:     item.CreateBy,
			UpdateTime:   common.TimeToString(item.UpdateTime),
			UpdateBy:     item.UpdateBy,
		})
	}

	return &umsclient.QueryMemberTaskListResp{
		Total: count,
		List:  list,
	}, nil
}
