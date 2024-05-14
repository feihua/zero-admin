package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionListByDateLogic 秒杀活动
/*
Author: LiuFeiHua
Date: 2023/12/5 13:53
*/
type FlashPromotionListByDateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionListByDateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListByDateLogic {
	return &FlashPromotionListByDateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionListByDate 查询当前时间是否有秒杀活动
func (l *FlashPromotionListByDateLogic) FlashPromotionListByDate(in *smsclient.FlashPromotionListByDateReq) (*smsclient.FlashPromotionListByDateResp, error) {
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

	logc.Infof(l.ctx, "查询限时购列表信息,参数：%+v,响应：%+v", in, list)

	return &smsclient.FlashPromotionListByDateResp{List: list}, nil
}
