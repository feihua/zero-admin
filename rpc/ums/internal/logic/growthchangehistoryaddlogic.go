package logic

import (
	"context"
	"time"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	// todo: add your logic here and delete this line
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsGrowthChangeHistoryModel.Insert(l.ctx, &umsmodel.UmsGrowthChangeHistory{
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
	return &umsclient.GrowthChangeHistoryAddResp{}, nil
}
