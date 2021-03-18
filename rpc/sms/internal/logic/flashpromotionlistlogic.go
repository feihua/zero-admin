package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListLogic {
	return &FlashPromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(in *sms.FlashPromotionListReq) (*sms.FlashPromotionListResp, error) {
	all, _ := l.svcCtx.SmsFlashPromotionModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionModel.Count()

	var list []*sms.FlashPromotionListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionListData{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:    item.EndDate.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &sms.FlashPromotionListResp{
		Total: count,
		List:  list,
	}, nil
}
