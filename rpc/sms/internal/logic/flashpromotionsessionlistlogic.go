package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionListLogic {
	return &FlashPromotionSessionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(in *sms.FlashPromotionSessionListReq) (*sms.FlashPromotionSessionListResp, error) {
	all, err := l.svcCtx.SmsFlashPromotionSessionModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionSessionModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购场次列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.FlashPromotionSessionListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionSessionListData{
			Id:         item.Id,
			Name:       item.Name,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询限时购场次列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sms.FlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil
}
