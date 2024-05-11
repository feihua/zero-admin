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
	_, err := l.svcCtx.GrowthChangeHistoryService.GrowthChangeHistoryDelete(l.ctx, &umsclient.GrowthChangeHistoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除成长值变化历史记录异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除成长值变化历史记录失败")
	}
	return &types.DeleteGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "删除成长值变化历史记录成功",
	}, nil
}
