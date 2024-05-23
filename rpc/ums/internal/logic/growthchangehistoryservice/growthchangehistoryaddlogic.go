package growthchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GrowthChangeHistoryAddLogic 成长值变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 9:11
*/
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

// GrowthChangeHistoryAdd 添加成长值变化历史记录
func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(in *umsclient.GrowthChangeHistoryAddReq) (*umsclient.GrowthChangeHistoryAddResp, error) {
	err := query.UmsGrowthChangeHistory.WithContext(l.ctx).Create(&model.UmsGrowthChangeHistory{
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
	return &umsclient.GrowthChangeHistoryAddResp{}, nil
}
