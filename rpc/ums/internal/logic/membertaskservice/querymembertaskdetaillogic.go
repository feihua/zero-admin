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

// QueryMemberTaskDetailLogic 查询会员任务表详情
/*
Author: LiuFeiHua
Date: 2024/6/11 13:41
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

// QueryMemberTaskDetail 查询会员任务表详情
func (l *QueryMemberTaskDetailLogic) QueryMemberTaskDetail(in *umsclient.QueryMemberTaskDetailReq) (*umsclient.QueryMemberTaskDetailResp, error) {
	item, err := query.UmsMemberTask.WithContext(l.ctx).Where(query.UmsMemberTask.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	data := &umsclient.QueryMemberTaskDetailResp{
		Id:           item.ID,
		TaskName:     item.TaskName,
		TaskGrowth:   item.TaskGrowth,
		TaskIntegral: item.TaskIntegral,
		TaskType:     item.TaskType,
		CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"),
		CreateBy:     item.CreateBy,
		UpdateTime:   common.TimeToString(item.UpdateTime),
		UpdateBy:     item.UpdateBy,
	}

	return data, nil
}
