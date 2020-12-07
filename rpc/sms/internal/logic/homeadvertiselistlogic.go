package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	all, _ := l.svcCtx.SmsHomeAdvertiseModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sms.HomeAdvertiseListData
	for _, item := range *all {

		list = append(list, &sms.HomeAdvertiseListData{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			Url:        item.Url,
			Note:       item.Note,
			Sort:       item.Sort,
		})
	}

	fmt.Println(list)
	return &sms.HomeAdvertiseListResp{
		Total: 10,
		List:  list,
	}, nil
}
