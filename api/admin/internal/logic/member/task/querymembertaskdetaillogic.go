package task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskDetailLogic 查询会员任务表详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type QueryMemberTaskDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskDetailLogic {
	return &QueryMemberTaskDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberTaskDetail 查询会员任务表详情
func (l *QueryMemberTaskDetailLogic) QueryMemberTaskDetail(req *types.QueryMemberTaskDetailReq) (resp *types.QueryMemberTaskDetailResp, err error) {

	detail, err := l.svcCtx.MemberTaskService.QueryMemberTaskDetail(l.ctx, &umsclient.QueryMemberTaskDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberTaskDetailData{
		Id:           detail.Id,           //
		TaskName:     detail.TaskName,     // 任务名称
		TaskGrowth:   detail.TaskGrowth,   // 赠送成长值
		TaskIntegral: detail.TaskIntegral, // 赠送积分
		TaskType:     detail.TaskType,     // 任务类型：0->新手任务；1->日常任务
		Status:       detail.Status,       // 状态：0->禁用；1->启用
		CreateBy:     detail.CreateBy,     // 创建者
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新者
		UpdateTime:   detail.UpdateTime,   // 更新时间
	}
	return &types.QueryMemberTaskDetailResp{
		Code:    "000000",
		Message: "查询会员任务表成功",
		Data:    data,
	}, nil
}
