package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.UmsIntegrationChangeHistoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsIntegrationChangeHistoryModel.Count()

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

	fmt.Println(list)
	return &ums.IntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
