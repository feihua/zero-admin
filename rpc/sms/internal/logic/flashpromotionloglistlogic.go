package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogListLogic {
	return &FlashPromotionLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionLogListLogic) FlashPromotionLogList(in *sms.FlashPromotionLogListReq) (*sms.FlashPromotionLogListResp, error) {
	all, _ := l.svcCtx.SmsFlashPromotionLogModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionLogModel.Count()

	var list []*sms.FlashPromotionLogListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionLogListData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			ProductId:     item.ProductId,
			MemberPhone:   item.MemberPhone,
			ProductName:   item.ProductName,
			SubscribeTime: item.SubscribeTime.Format("2006-01-02 15:04:05"),
			SendTime:      item.SendTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &sms.FlashPromotionLogListResp{
		Total: count,
		List:  list,
	}, nil
}
