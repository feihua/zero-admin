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
	_, err := l.svcCtx.IntegrationChangeHistoryService.IntegrationChangeHistoryDelete(l.ctx, &umsclient.IntegrationChangeHistoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除积分变化历史记录异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除积分变化历史记录失败")
	}
	return &types.DeleteIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "查询成长值变化历史记录成功",
	}, nil
}
