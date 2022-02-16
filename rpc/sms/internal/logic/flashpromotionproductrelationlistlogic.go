package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
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
	all, err := l.svcCtx.SmsFlashPromotionProductRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionProductRelationModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购与产品关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询限时购与产品关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sms.FlashPromotionProductRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
