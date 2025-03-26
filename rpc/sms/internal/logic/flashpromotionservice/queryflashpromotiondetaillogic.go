package flashpromotionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionDetailLogic 查询限时购详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
*/
type QueryFlashPromotionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionDetailLogic {
	return &QueryFlashPromotionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionDetail 查询限时购详情
func (l *QueryFlashPromotionDetailLogic) QueryFlashPromotionDetail(in *smsclient.QueryFlashPromotionDetailReq) (*smsclient.QueryFlashPromotionDetailResp, error) {
	item, err := query.SmsFlashPromotion.WithContext(l.ctx).Where(query.SmsFlashPromotion.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询限时购详情失败")
	}

	data := &smsclient.QueryFlashPromotionDetailResp{
		Id:         item.ID,                              // 编号
		Title:      item.Title,                           // 标题
		StartDate:  item.StartDate.Format("2006-01-02"),  // 开始日期
		EndDate:    item.EndDate.Format("2006-01-02"),    // 结束日期
		Status:     item.Status,                          // 上下线状态
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
	}

	return data, nil
}
