package growthchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddGrowthChangeHistoryLogic 添加成长值变化历史记录表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:24
*/
type AddGrowthChangeHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGrowthChangeHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGrowthChangeHistoryLogic {
	return &AddGrowthChangeHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddGrowthChangeHistory 添加成长值变化历史记录表
func (l *AddGrowthChangeHistoryLogic) AddGrowthChangeHistory(in *umsclient.AddGrowthChangeHistoryReq) (*umsclient.AddGrowthChangeHistoryResp, error) {
	err := query.UmsGrowthChangeHistory.WithContext(l.ctx).Create(&model.UmsGrowthChangeHistory{
		MemberID:    in.MemberId,
		CreateTime:  time.Now(),
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: in.OperateNote,
		SourceType:  in.SourceType,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddGrowthChangeHistoryResp{}, nil
}
