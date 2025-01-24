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
)

// AddMemberTaskLogic 添加会员任务
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type AddMemberTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTaskLogic {
	return &AddMemberTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTask 添加会员任务
func (l *AddMemberTaskLogic) AddMemberTask(in *umsclient.AddMemberTaskReq) (*umsclient.AddMemberTaskResp, error) {
	q := query.UmsMemberTask

	item := &model.UmsMemberTask{
		TaskName:     in.TaskName,     // 任务名称
		TaskGrowth:   in.TaskGrowth,   // 赠送成长值
		TaskIntegral: in.TaskIntegral, // 赠送积分
		TaskType:     in.TaskType,     // 任务类型：0->新手任务；1->日常任务
		Status:       in.Status,       // 状态：0->禁用；1->启用
		CreateBy:     in.CreateBy,     // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员任务失败")
	}

	logc.Infof(l.ctx, "添加会员任务成功,参数：%+v", in)
	return &umsclient.AddMemberTaskResp{}, nil
}
