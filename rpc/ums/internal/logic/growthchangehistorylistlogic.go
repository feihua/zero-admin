package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *GrowthChangeHistoryListLogic) GrowthChangeHistoryList(in *ums.GrowthChangeHistoryListReq) (*ums.GrowthChangeHistoryListResp, error) {
	all, _ := l.svcCtx.UmsGrowthChangeHistoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsGrowthChangeHistoryModel.Count()

	var list []*ums.GrowthChangeHistoryListData
	for _, item := range *all {

		list = append(list, &ums.GrowthChangeHistoryListData{
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
	return &ums.GrowthChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
