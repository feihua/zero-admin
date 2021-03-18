package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionListLogic {
	return &FlashPromotionSessionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(in *sms.FlashPromotionSessionListReq) (*sms.FlashPromotionSessionListResp, error) {
	all, _ := l.svcCtx.SmsFlashPromotionSessionModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionSessionModel.Count()

	var list []*sms.FlashPromotionSessionListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionSessionListData{
			Id:         item.Id,
			Name:       item.Name,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &sms.FlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil
}
