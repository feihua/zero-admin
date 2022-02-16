package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationChangeHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationChangeHistoryListLogic {
	return &IntegrationChangeHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(in *ums.IntegrationChangeHistoryListReq) (*ums.IntegrationChangeHistoryListResp, error) {
	all, err := l.svcCtx.UmsIntegrationChangeHistoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsIntegrationChangeHistoryModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分变化历史记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.IntegrationChangeHistoryListData
	for _, item := range *all {

		list = append(list, &ums.IntegrationChangeHistoryListData{
			Id:          item.Id,
			MemberId:    item.MemberId,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			ChangeType:  item.ChangeType,
			ChangeCount: item.ChangeCount,
			OperateMan:  item.OperateMan,
			OperateNote: item.OperateNote,
			SourceType:  item.SourceType,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询积分变化历史记录列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.IntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
