package integrationchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.UmsIntegrationChangeHistory.WithContext(l.ctx)
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询积分变化历史记录列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.IntegrationChangeHistoryListData
	for _, item := range result {

		list = append(list, &umsclient.IntegrationChangeHistoryListData{
			Id:          item.ID,
			MemberId:    item.MemberID,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			ChangeType:  item.ChangeType,
			ChangeCount: item.ChangeCount,
			OperateMan:  item.OperateMan,
			OperateNote: *item.OperateNote,
			SourceType:  item.SourceType,
		})
	}

	logc.Infof(l.ctx, "查询积分变化历史记录列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.IntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
