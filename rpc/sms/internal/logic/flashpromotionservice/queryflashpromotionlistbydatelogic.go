package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionListByDateLogic 查询当前时间是否有秒杀活动
/*
Author: LiuFeiHua
Date: 2024/6/12 17:47
*/
type QueryFlashPromotionListByDateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionListByDateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionListByDateLogic {
	return &QueryFlashPromotionListByDateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionListByDate 查询当前时间是否有秒杀活动
func (l *QueryFlashPromotionListByDateLogic) QueryFlashPromotionListByDate(in *smsclient.QueryFlashPromotionListByDateReq) (*smsclient.QueryFlashPromotionListByDateResp, error) {
	currentDate, _ := time.Parse("2006-01-02", in.CurrentDate)
	q := query.SmsFlashPromotion
	result, err := q.WithContext(l.ctx).Where(q.Status.Eq(1), q.StartDate.Lte(currentDate), q.EndDate.Gte(currentDate)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionListData{
			Id:         item.ID,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02"),
			EndDate:    item.EndDate.Format("2006-01-02"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &smsclient.QueryFlashPromotionListByDateResp{List: list}, nil

}
