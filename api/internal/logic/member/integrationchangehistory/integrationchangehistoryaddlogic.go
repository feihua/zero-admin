package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员积分变化历史记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加积分变化历史记录失败")
	}

	return &types.AddIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "添加积分变化历史记录成功",
	}, nil
}
