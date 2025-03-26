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

// QueryMemberTaskDetailLogic 查询会员任务详情
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
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

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员任务详情失败")
	}

	data := &umsclient.QueryMemberTaskDetailResp{
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
	}

	return data, nil
}
