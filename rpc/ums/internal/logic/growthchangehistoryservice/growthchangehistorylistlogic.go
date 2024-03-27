package growthchangehistoryservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(in *umsclient.GrowthChangeHistoryListReq) (*umsclient.GrowthChangeHistoryListResp, error) {
	all, err := l.svcCtx.UmsGrowthChangeHistoryModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsGrowthChangeHistoryModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询成长值变化历史记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.GrowthChangeHistoryListData
	for _, item := range *all {

		list = append(list, &umsclient.GrowthChangeHistoryListData{
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
	logx.WithContext(l.ctx).Infof("查询成长值变化历史记录列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.GrowthChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
