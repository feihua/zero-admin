package integrationchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationChangeHistoryAddLogic 积分变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 9:57
*/
type IntegrationChangeHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryAddLogic {
	return &IntegrationChangeHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationChangeHistoryAdd 添加积分变化历史记录
func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(in *umsclient.IntegrationChangeHistoryAddReq) (*umsclient.IntegrationChangeHistoryAddResp, error) {
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

	return &umsclient.IntegrationChangeHistoryAddResp{}, nil
}
