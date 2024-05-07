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

// IntegrationChangeHistoryUpdateLogic 积分变化历史记录
type IntegrationChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryUpdateLogic {
	return &IntegrationChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IntegrationChangeHistoryUpdate 更新积分变化历史记录
func (l *IntegrationChangeHistoryUpdateLogic) IntegrationChangeHistoryUpdate(in *umsclient.IntegrationChangeHistoryUpdateReq) (*umsclient.IntegrationChangeHistoryUpdateResp, error) {
	_, err := query.UmsIntegrationChangeHistory.WithContext(l.ctx).Updates(&model.UmsIntegrationChangeHistory{
		ID:          in.Id,
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

	return &umsclient.IntegrationChangeHistoryUpdateResp{}, nil
}
