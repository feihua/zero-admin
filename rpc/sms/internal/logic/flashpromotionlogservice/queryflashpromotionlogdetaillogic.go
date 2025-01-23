package flashpromotionlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionLogDetailLogic 查询限时购通知记录详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
*/
type QueryFlashPromotionLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionLogDetailLogic {
	return &QueryFlashPromotionLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionLogDetail 查询限时购通知记录详情
func (l *QueryFlashPromotionLogDetailLogic) QueryFlashPromotionLogDetail(in *smsclient.QueryFlashPromotionLogDetailReq) (*smsclient.QueryFlashPromotionLogDetailResp, error) {
	item, err := query.SmsFlashPromotionLog.WithContext(l.ctx).Where(query.SmsFlashPromotionLog.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购通知记录详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询限时购通知记录详情失败")
	}

	data := &smsclient.QueryFlashPromotionLogDetailResp{
		Id:            item.ID,                                          // 编号
		MemberId:      item.MemberID,                                    // 会员id
		ProductId:     item.ProductID,                                   // 商品id
		MemberPhone:   item.MemberPhone,                                 // 会员电话
		ProductName:   item.ProductName,                                 // 商品名称
		SubscribeTime: item.SubscribeTime.Format("2006-01-02 15:04:05"), // 会员订阅时间
		SendTime:      item.SendTime.Format("2006-01-02 15:04:05"),      // 发送时间
	}

	logc.Infof(l.ctx, "查询限时购通知记录详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
