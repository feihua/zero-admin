package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionProductRelationListLogic 查询商品限时购与商品关系表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:41
*/
type QueryFlashPromotionProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionProductRelationListLogic {
	return &QueryFlashPromotionProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionProductRelationList 查询商品限时购与商品关系表列表
func (l *QueryFlashPromotionProductRelationListLogic) QueryFlashPromotionProductRelationList(in *smsclient.QueryFlashPromotionProductRelationListReq) (*smsclient.QueryFlashPromotionProductRelationListResp, error) {
	relation := query.SmsFlashPromotionProductRelation
	q := relation.WithContext(l.ctx)
	q = q.Where(relation.FlashPromotionID.Eq(in.FlashPromotionId), relation.FlashPromotionSessionID.Eq(in.FlashPromotionSessionId))

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购与产品关糸列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionProductRelationListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionProductRelationListData{
			Id:                      item.ID,
			FlashPromotionId:        item.FlashPromotionID,
			FlashPromotionSessionId: item.FlashPromotionSessionID,
			ProductId:               item.ProductID,
			FlashPromotionPrice:     item.FlashPromotionPrice,
			FlashPromotionCount:     item.FlashPromotionCount,
			FlashPromotionLimit:     item.FlashPromotionLimit,
			Sort:                    item.Sort,
		})
	}

	return &smsclient.QueryFlashPromotionProductRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
