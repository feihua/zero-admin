package flashpromotionlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionLogListLogic 限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/5/14 10:53
*/
type FlashPromotionLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionLogListLogic {
	return &FlashPromotionLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionLogList 查询限时购通知记录
func (l *FlashPromotionLogListLogic) FlashPromotionLogList(in *smsclient.FlashPromotionLogListReq) (*smsclient.FlashPromotionLogListResp, error) {
	q := query.SmsFlashPromotionLog.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购通知记录列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionLogListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionLogListData{
			Id:            item.ID,
			MemberId:      item.MemberID,
			ProductId:     item.ProductID,
			MemberPhone:   item.MemberPhone,
			ProductName:   item.ProductName,
			SubscribeTime: item.SubscribeTime.Format("2006-01-02 15:04:05"),
			SendTime:      item.SendTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询限时购通知记录列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.FlashPromotionLogListResp{
		Total: count,
		List:  list,
	}, nil
}
