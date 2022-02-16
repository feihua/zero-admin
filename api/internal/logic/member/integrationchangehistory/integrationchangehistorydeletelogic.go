package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryDeleteLogic {
	return IntegrationChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(req types.DeleteIntegrationChangeHistoryReq) (*types.DeleteIntegrationChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.IntegrationChangeHistoryDelete(l.ctx, &umsclient.IntegrationChangeHistoryDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除积分变化历史记录异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除积分变化历史记录失败")
	}
	return &types.DeleteIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "查询成长值变化历史记录成功",
	}, nil
}
