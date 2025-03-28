package flashpromotionsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionSessionListByTimeLogic 根据时间查询限时购场次
/*
Author: LiuFeiHua
Date: 2024/6/12 17:49
*/
type QueryFlashPromotionSessionListByTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionSessionListByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionSessionListByTimeLogic {
	return &QueryFlashPromotionSessionListByTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionSessionListByTime 根据时间查询限时购场次
func (l *QueryFlashPromotionSessionListByTimeLogic) QueryFlashPromotionSessionListByTime(in *smsclient.QueryFlashPromotionSessionListByTimeReq) (*smsclient.QueryFlashPromotionSessionListByTimeResp, error) {
	// times := strings.Split(in.CurrentTIme, " ")[1]
	// currentTIme, _ := time_util.Parse("15:04:05", times)
	q := query.SmsFlashPromotionSession
	// result, err := q.WithContext(l.ctx).Where(q.Status.Eq(1), q.StartTime.Lte(currentTIme), q.EndTime.Gte(currentTIme)).Find()
	result, err := q.WithContext(l.ctx).Where(q.Status.Eq(1)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "根据时间查询限时购场次失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据时间查询限时购场次失败")
	}

	var list []*smsclient.FlashPromotionSessionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionSessionListData{
			Id:         item.ID,                              // 编号
			Name:       item.Name,                            // 场次名称
			StartTime:  item.StartTime,                       // 每日开始时间
			EndTime:    item.EndTime,                         // 每日结束时间
			Status:     item.Status,                          // 启用状态：0->不启用；1->启用
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
		})
	}

	return &smsclient.QueryFlashPromotionSessionListByTimeResp{
		List: list,
	}, nil
}
