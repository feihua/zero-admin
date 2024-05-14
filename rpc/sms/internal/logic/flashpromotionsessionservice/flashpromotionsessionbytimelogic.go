package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionByTimeLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:47
*/
type FlashPromotionSessionByTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionByTimeLogic {
	return &FlashPromotionSessionByTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionSessionByTime 根据时间查询限时购场次
func (l *FlashPromotionSessionByTimeLogic) FlashPromotionSessionByTime(in *smsclient.FlashPromotionSessionByTimeReq) (*smsclient.FlashPromotionSessionByTimeResp, error) {
	times := strings.Split(in.CurrentTIme, " ")[1]
	currentTIme, _ := time.Parse("15:04:05", times)
	q := query.SmsFlashPromotionSession
	result, err := q.WithContext(l.ctx).Where(q.Status.Eq(1), q.StartTime.Lte(currentTIme), q.EndTime.Gte(currentTIme)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购场次列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionSessionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionSessionListData{
			Id:         item.ID,
			Name:       item.Name,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询限时购场次列表信息,参数：%+v,响应：%+v", in, list)

	return &smsclient.FlashPromotionSessionByTimeResp{}, nil
}
