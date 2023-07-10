package integrationchangehistoryservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(in *umsclient.IntegrationChangeHistoryAddReq) (*umsclient.IntegrationChangeHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsIntegrationChangeHistoryModel.Insert(l.ctx, &umsmodel.UmsIntegrationChangeHistory{
		MemberId:    in.MemberId,
		CreateTime:  CreateTime,
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: sql.NullString{String: in.OperateNote, Valid: true},
		SourceType:  in.SourceType,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.IntegrationChangeHistoryAddResp{}, nil
}
