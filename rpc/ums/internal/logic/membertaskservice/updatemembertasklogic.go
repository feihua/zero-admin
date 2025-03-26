package membertaskservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTaskLogic 更新会员任务
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
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
	_, err := q.Where(query.UmsMemberTask.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据会员任务id：%d,查询会员任务失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询会员任务失败"))
	}

	item := &model.UmsMemberTask{
		ID:           in.Id,           //
		TaskName:     in.TaskName,     // 任务名称
		TaskGrowth:   in.TaskGrowth,   // 赠送成长值
		TaskIntegral: in.TaskIntegral, // 赠送积分
		TaskType:     in.TaskType,     // 任务类型：0->新手任务；1->日常任务
		Status:       in.Status,       // 状态：0->禁用；1->启用
		UpdateBy:     in.UpdateBy,     // 更新者
	}

	// 2.会员任务存在时,则直接更新会员任务
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新会员任务失败")
	}

	return &umsclient.UpdateMemberTaskResp{}, nil
}
