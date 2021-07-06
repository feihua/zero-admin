package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryAddLogic {
	return IntegrationChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(req types.AddIntegrationChangeHistoryReq) (*types.AddIntegrationChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.IntegrationChangeHistoryAdd(l.ctx, &umsclient.IntegrationChangeHistoryAddReq{
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加积分变化历史记录失败")
	}

	return &types.AddIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "添加积分变化历史记录成功",
	}, nil
}
