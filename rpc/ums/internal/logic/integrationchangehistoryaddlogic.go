package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *IntegrationChangeHistoryAddLogic) IntegrationChangeHistoryAdd(in *ums.IntegrationChangeHistoryAddReq) (*ums.IntegrationChangeHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsIntegrationChangeHistoryModel.Insert(umsmodel.UmsIntegrationChangeHistory{
		MemberId:    in.MemberId,
		CreateTime:  CreateTime,
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: in.OperateNote,
		SourceType:  in.SourceType,
	})
	if err != nil {
		return nil, err
	}

	return &ums.IntegrationChangeHistoryAddResp{}, nil
}
