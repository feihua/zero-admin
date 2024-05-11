package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.GrowthChangeHistoryService.GrowthChangeHistoryUpdate(l.ctx, &umsclient.GrowthChangeHistoryUpdateReq{
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
		logc.Errorf(l.ctx, "更新成长值变化历史记录信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新成长值变化历史记录失败")
	}

	return &types.UpdateGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "更新成长值变化历史记录成功",
	}, nil
}
