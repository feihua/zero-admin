package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskAddLogic {
	return MemberTaskAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(req types.AddMemberTaskReq) (*types.AddMemberTaskResp, error) {
	_, err := l.svcCtx.MemberTaskService.MemberTaskAdd(l.ctx, &umsclient.MemberTaskAddReq{
		TaskName:     req.Name,
		TaskGrowth:   req.Growth,
		TaskIntegral: req.Intergration,
		TaskType:     req.Type,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员任务信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员任务失败")
	}

	return &types.AddMemberTaskResp{
		Code:    "000000",
		Message: "添加会员任务成功",
	}, nil
}
