package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryUpdateLogic {
	return &GrowthChangeHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(in *ums.GrowthChangeHistoryUpdateReq) (*ums.GrowthChangeHistoryUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.UmsGrowthChangeHistoryModel.Update(l.ctx, &umsmodel.UmsGrowthChangeHistory{
		Id:          in.Id,
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

	return &ums.GrowthChangeHistoryUpdateResp{}, nil
}
