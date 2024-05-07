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

// GrowthChangeHistoryUpdateLogic 成长值变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 9:54
*/
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

// GrowthChangeHistoryUpdate 添加成长值变化历史记录
func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(in *umsclient.GrowthChangeHistoryUpdateReq) (*umsclient.GrowthChangeHistoryUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := query.UmsGrowthChangeHistory.WithContext(l.ctx).Updates(&model.UmsGrowthChangeHistory{
		ID:          in.Id,
		MemberID:    in.MemberId,
		CreateTime:  CreateTime,
		ChangeType:  in.ChangeType,
		ChangeCount: in.ChangeCount,
		OperateMan:  in.OperateMan,
		OperateNote: &in.OperateNote,
		SourceType:  in.SourceType,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.GrowthChangeHistoryUpdateResp{}, nil
}
