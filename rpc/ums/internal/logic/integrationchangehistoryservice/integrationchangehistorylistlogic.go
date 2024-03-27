package integrationchangehistoryservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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

func (l *IntegrationChangeHistoryListLogic) IntegrationChangeHistoryList(in *umsclient.IntegrationChangeHistoryListReq) (*umsclient.IntegrationChangeHistoryListResp, error) {
	all, err := l.svcCtx.UmsIntegrationChangeHistoryModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsIntegrationChangeHistoryModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分变化历史记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.IntegrationChangeHistoryListData
	for _, item := range *all {

		list = append(list, &umsclient.IntegrationChangeHistoryListData{
			Id:          item.Id,
			MemberId:    item.MemberId,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			ChangeType:  item.ChangeType,
			ChangeCount: item.ChangeCount,
			OperateMan:  item.OperateMan,
			OperateNote: item.OperateNote.String,
			SourceType:  item.SourceType,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询积分变化历史记录列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.IntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
