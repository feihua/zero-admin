package integrationchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationChangeHistoryListLogic 查询积分变化历史记录表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:24
*/
type QueryIntegrationChangeHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIntegrationChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationChangeHistoryListLogic {
	return &QueryIntegrationChangeHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryIntegrationChangeHistoryList 查询积分变化历史记录表列表
func (l *QueryIntegrationChangeHistoryListLogic) QueryIntegrationChangeHistoryList(in *umsclient.QueryIntegrationChangeHistoryListReq) (*umsclient.QueryIntegrationChangeHistoryListResp, error) {
	history := query.UmsIntegrationChangeHistory
	q := history.WithContext(l.ctx).Where(history.MemberID.Eq(in.MemberId))
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

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
			OperateNote: item.OperateNote,
			SourceType:  item.SourceType,
		})
	}

	return &umsclient.QueryIntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
