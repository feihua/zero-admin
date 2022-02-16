package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrowthChangeHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryUpdateLogic {
	return GrowthChangeHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryUpdateLogic) GrowthChangeHistoryUpdate(req types.UpdateGrowthChangeHistoryReq) (*types.UpdateGrowthChangeHistoryResp, error) {
	_, err := l.svcCtx.Ums.GrowthChangeHistoryUpdate(l.ctx, &umsclient.GrowthChangeHistoryUpdateReq{
		Id:          req.Id,
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新成长值变化历史记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新成长值变化历史记录失败")
	}

	return &types.UpdateGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "更新成长值变化历史记录成功",
	}, nil
}
