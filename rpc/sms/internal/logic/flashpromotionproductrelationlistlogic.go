package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationListLogic {
	return &FlashPromotionProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationListLogic) FlashPromotionProductRelationList(in *sms.FlashPromotionProductRelationListReq) (*sms.FlashPromotionProductRelationListResp, error) {
	all, _ := l.svcCtx.SmsFlashPromotionProductRelationModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sms.FlashPromotionProductRelationListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionProductRelationListData{
			Id:                      item.Id,
			FlashPromotionId:        item.FlashPromotionId,
			FlashPromotionSessionId: item.FlashPromotionSessionId,
			ProductId:               item.ProductId,
			FlashPromotionPrice:     int64(item.FlashPromotionPrice),
			FlashPromotionCount:     item.FlashPromotionCount,
			FlashPromotionLimit:     item.FlashPromotionLimit,
			Sort:                    item.Sort,
		})
	}

	fmt.Println(list)
	return &sms.FlashPromotionProductRelationListResp{
		Total: 10,
		List:  list,
	}, nil
}
