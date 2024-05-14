package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	relation := query.SmsFlashPromotionProductRelation
	q := relation.WithContext(l.ctx)
	q = q.Where(relation.FlashPromotionID.Eq(in.FlashPromotionId), relation.FlashPromotionSessionID.Eq(in.FlashPromotionSessionId))

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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

	logc.Infof(l.ctx, "查询限时购与产品关糸列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.FlashPromotionProductRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
