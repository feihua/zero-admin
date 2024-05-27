package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTaskListLogic 查询会员任务列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:52
*/
type MemberTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskListLogic {
	return &MemberTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTaskList 查询会员任务列表
func (l *MemberTaskListLogic) MemberTaskList(in *umsclient.MemberTaskListReq) (*umsclient.MemberTaskListResp, error) {
	q := query.UmsMemberTask.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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

	logc.Infof(l.ctx, "查询会员任务列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberTaskListResp{
		Total: count,
		List:  list,
	}, nil

}
