package growthchangehistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// GrowthChangeHistoryListLogic 成长值变化历史记录
/*
Author: LiuFeiHua
Date: 2024/5/7 10:50
*/
type GrowthChangeHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrowthChangeHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrowthChangeHistoryListLogic {
	return &GrowthChangeHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GrowthChangeHistoryList 查询成长值变化历史记录
func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(in *umsclient.GrowthChangeHistoryListReq) (*umsclient.GrowthChangeHistoryListResp, error) {
	history := query.UmsGrowthChangeHistory
	q := history.WithContext(l.ctx).Where(history.MemberID.Eq(in.MemberId))
	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询成长值变化历史记录列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.GrowthChangeHistoryListData
	for _, item := range result {

		list = append(list, &umsclient.GrowthChangeHistoryListData{
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

	logc.Infof(l.ctx, "查询成长值变化历史记录列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.GrowthChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
