package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryGrowthChangeHistoryListLogic 查询成长值变化历史记录列表
/*
Author: LiuFeiHua
Date: 2024/5/22 9:31
*/
type QueryGrowthChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGrowthChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGrowthChangeHistoryListLogic {
	return &QueryGrowthChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryGrowthChangeHistoryList 查询成长值变化历史记录列表
func (l *QueryGrowthChangeHistoryListLogic) QueryGrowthChangeHistoryList(req *types.ListChangeHistoryReq) (resp *types.ListChangeHistoryResp, err error) {
	result, err := l.svcCtx.GrowthChangeHistoryService.GrowthChangeHistoryList(l.ctx, &umsclient.GrowthChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询成长值变化历史记录列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询成长值变化历史记录失败")
	}

	var list []*types.ListChangeHistoryData

	for _, item := range result.List {
		list = append(list, &types.ListChangeHistoryData{
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

	return &types.ListChangeHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询成长值变化历史记录成功",
	}, nil
}
