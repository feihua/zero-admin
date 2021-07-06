package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/ums/umsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryAddLogic {
	return GrowthChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(req types.AddGrowthChangeHistoryReq) (*types.AddGrowthChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.GrowthChangeHistoryAdd(l.ctx, &umsclient.GrowthChangeHistoryAddReq{
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加成长值变化历史记录失败")
	}

	return &types.AddGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "添加成长值变化历史记录成功",
	}, nil
}
