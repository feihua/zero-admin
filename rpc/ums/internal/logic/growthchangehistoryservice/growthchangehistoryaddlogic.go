package growthchangehistoryservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryAddLogic {
	return &GrowthChangeHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(in *umsclient.GrowthChangeHistoryAddReq) (*umsclient.GrowthChangeHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsGrowthChangeHistoryModel.Insert(l.ctx, &umsmodel.UmsGrowthChangeHistory{
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

	return &umsclient.GrowthChangeHistoryAddResp{}, nil
}
