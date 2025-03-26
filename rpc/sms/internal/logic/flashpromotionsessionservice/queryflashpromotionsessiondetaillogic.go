package flashpromotionsessionservicelogic

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

// QueryFlashPromotionSessionDetailLogic 查询限时购场次详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
*/
type QueryFlashPromotionSessionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionSessionDetailLogic {
	return &QueryFlashPromotionSessionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionSessionDetail 查询限时购场次详情
func (l *QueryFlashPromotionSessionDetailLogic) QueryFlashPromotionSessionDetail(in *smsclient.QueryFlashPromotionSessionDetailReq) (*smsclient.QueryFlashPromotionSessionDetailResp, error) {
	item, err := query.SmsFlashPromotionSession.WithContext(l.ctx).Where(query.SmsFlashPromotionSession.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购场次详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询限时购场次详情失败")
	}

	data := &smsclient.QueryFlashPromotionSessionDetailResp{
		Id:         item.ID,                              // 编号
		Name:       item.Name,                            // 场次名称
		StartTime:  item.StartTime,                       // 每日开始时间
		EndTime:    item.EndTime,                         // 每日结束时间
		Status:     item.Status,                          // 启用状态：0->不启用；1->启用
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
	}

	return data, nil
}
