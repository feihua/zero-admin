package flashpromotionlogservicelogic

import (
	"context"
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
	return &smsclient.QueryFlashPromotionLogListResp{
		Total: count,
		List:  list,
	}, nil
}
