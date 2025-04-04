package flashpromotionlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionLogListLogic 查询限时购通知记录列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:39
*/
type QueryFlashPromotionLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionLogListLogic {
	return &QueryFlashPromotionLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionLogList 查询限时购通知记录列表
func (l *QueryFlashPromotionLogListLogic) QueryFlashPromotionLogList(in *smsclient.QueryFlashPromotionLogListReq) (*smsclient.QueryFlashPromotionLogListResp, error) {
	q := query.SmsFlashPromotionLog.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购通知记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询限时购通知记录列表失败")
	}

	var list []*smsclient.FlashPromotionLogListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionLogListData{
			Id:            item.ID,                                          // 编号
			MemberId:      item.MemberID,                                    // 会员id
			ProductId:     item.ProductID,                                   // 商品id
			MemberPhone:   item.MemberPhone,                                 // 会员电话
			ProductName:   item.ProductName,                                 // 商品名称
			SubscribeTime: item.SubscribeTime.Format("2006-01-02 15:04:05"), // 会员订阅时间
			SendTime:      item.SendTime.Format("2006-01-02 15:04:05"),      // 发送时间
		})
	}

	return &smsclient.QueryFlashPromotionLogListResp{
		Total: count,
		List:  list,
	}, nil
}
