package integrationchangehistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationChangeHistoryListLogic 查询积分变化历史记录列表
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

// QueryIntegrationChangeHistoryList 查询积分变化历史记录列表
func (l *QueryIntegrationChangeHistoryListLogic) QueryIntegrationChangeHistoryList(in *umsclient.QueryIntegrationChangeHistoryListReq) (*umsclient.QueryIntegrationChangeHistoryListResp, error) {
	history := query.UmsIntegrationChangeHistory
	q := history.WithContext(l.ctx).Where(history.MemberID.Eq(in.MemberId))
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分变化历史记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询积分变化历史记录列表失败")
	}

	var list []*umsclient.IntegrationChangeHistoryListData
	for _, item := range result {

		list = append(list, &umsclient.IntegrationChangeHistoryListData{
			Id:          item.ID,                              //
			MemberId:    item.MemberID,                        // 会员id
			ChangeType:  item.ChangeType,                      // 改变类型：0->增加；1->减少
			ChangeCount: item.ChangeCount,                     // 积分改变数量
			OperateMan:  item.OperateMan,                      // 操作人员
			OperateNote: item.OperateNote,                     // 操作备注
			SourceType:  item.SourceType,                      // 积分来源：0->购物；1->管理员修改
			CreateTime:  time_util.TimeToStr(item.CreateTime), // 创建时间
		})
	}

	return &umsclient.QueryIntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
