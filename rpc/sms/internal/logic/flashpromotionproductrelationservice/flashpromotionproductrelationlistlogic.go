package flashpromotionproductrelationservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

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

func (l *FlashPromotionProductRelationListLogic) FlashPromotionProductRelationList(in *smsclient.FlashPromotionProductRelationListReq) (*smsclient.FlashPromotionProductRelationListResp, error) {
	all, err := l.svcCtx.SmsFlashPromotionProductRelationModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SmsFlashPromotionProductRelationModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购与产品关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionProductRelationListData
	for _, item := range *all {

		list = append(list, &smsclient.FlashPromotionProductRelationListData{
			Id:                      item.Id,
			FlashPromotionId:        item.FlashPromotionId,
			FlashPromotionSessionId: item.FlashPromotionSessionId,
			ProductId:               item.ProductId,
			FlashPromotionPrice:     item.FlashPromotionPrice,
			FlashPromotionCount:     item.FlashPromotionCount,
			FlashPromotionLimit:     item.FlashPromotionLimit,
			Sort:                    item.Sort,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询限时购与产品关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.FlashPromotionProductRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
