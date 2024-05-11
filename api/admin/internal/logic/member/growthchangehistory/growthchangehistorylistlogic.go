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

type GrowthChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryListLogic {
	return GrowthChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(req types.ListGrowthChangeHistoryReq) (*types.ListGrowthChangeHistoryResp, error) {
	resp, err := l.svcCtx.GrowthChangeHistoryService.GrowthChangeHistoryList(l.ctx, &umsclient.GrowthChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询成长值变化历史记录列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询成长值变化历史记录失败")
	}

	var list []*types.ListtGrowthChangeHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtGrowthChangeHistoryData{
			Id:          item.Id,
			MemberId:    item.MemberId,
			CreateTime:  item.CreateTime,
			ChangeType:  item.ChangeType,
			ChangeCount: item.ChangeCount,
			OperateMan:  item.OperateMan,
			OperateNote: item.OperateNote,
			SourceType:  item.SourceType,
		})
	}

	return &types.ListGrowthChangeHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询成长值变化历史记录成功",
	}, nil
}
