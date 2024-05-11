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

type MemberTaskUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskUpdateLogic {
	return MemberTaskUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(req types.UpdateMemberTaskReq) (*types.UpdateMemberTaskResp, error) {
	_, err := l.svcCtx.MemberTaskService.MemberTaskUpdate(l.ctx, &umsclient.MemberTaskUpdateReq{
		Id:           req.Id,
		TaskName:     req.Name,
		TaskGrowth:   req.Growth,
		TaskIntegral: req.Intergration,
		TaskType:     req.Type,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员任务信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员任务失败")
	}

	return &types.UpdateMemberTaskResp{
		Code:    "000000",
		Message: "更新会员任务成功",
	}, nil
}
