package flashpromotionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
		logc.Errorf(l.ctx, "查询当前时间是否有秒杀活动失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询当前时间是否有秒杀活动失败")
	}

	var list []*smsclient.FlashPromotionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionListData{
			Id:         item.ID,                              // 编号
			Title:      item.Title,                           // 标题
			StartDate:  item.StartDate.Format("2006-01-02"),  // 开始日期
			EndDate:    item.EndDate.Format("2006-01-02"),    // 结束日期
			Status:     item.Status,                          // 上下线状态
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
		})
	}

	return &smsclient.QueryFlashPromotionListByDateResp{List: list}, nil

}
