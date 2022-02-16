package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryListLogic {
	return IntegrationChangeHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(req types.ListIntegrationChangeHistoryReq) (*types.ListIntegrationChangeHistoryResp, error) {
	resp, err := l.svcCtx.Ums.IntegrationChangeHistoryList(l.ctx, &umsclient.IntegrationChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询积分变化历史记录列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询积分变化历史记录失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtIntegrationChangeHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtIntegrationChangeHistoryData{
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

	return &types.ListIntegrationChangeHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
