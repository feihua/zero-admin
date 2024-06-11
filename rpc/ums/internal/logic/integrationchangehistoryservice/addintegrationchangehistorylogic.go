package integrationchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddIntegrationChangeHistoryLogic 添加积分变化历史记录表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:23
*/
type AddIntegrationChangeHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegrationChangeHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegrationChangeHistoryLogic {
	return &AddIntegrationChangeHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddIntegrationChangeHistory 添加积分变化历史记录表
func (l *AddIntegrationChangeHistoryLogic) AddIntegrationChangeHistory(in *umsclient.AddIntegrationChangeHistoryReq) (*umsclient.AddIntegrationChangeHistoryResp, error) {
	err := query.UmsIntegrationChangeHistory.WithContext(l.ctx).Create(&model.UmsIntegrationChangeHistory{
		MemberID:    in.MemberId,
		CreateTime:  time.Now(),
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: &in.OperateNote,
		SourceType:  in.SourceType,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddIntegrationChangeHistoryResp{}, nil
}
