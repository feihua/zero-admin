package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryDeleteLogic {
	return GrowthChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryDeleteLogic) GrowthChangeHistoryDelete(req types.DeleteGrowthChangeHistoryReq) (*types.DeleteGrowthChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.GrowthChangeHistoryDelete(l.ctx, &umsclient.GrowthChangeHistoryDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除成长值变化历史记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除成长值变化历史记录失败")
	}
	return &types.DeleteGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "删除成长值变化历史记录成功",
	}, nil
}
