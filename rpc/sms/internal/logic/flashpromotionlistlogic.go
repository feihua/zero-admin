package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListLogic {
	return &FlashPromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(in *sms.FlashPromotionListReq) (*sms.FlashPromotionListResp, error) {
	all, err := l.svcCtx.SmsFlashPromotionModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.FlashPromotionListData
	for _, item := range *all {

		list = append(list, &sms.FlashPromotionListData{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:    item.EndDate.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询限时购列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sms.FlashPromotionListResp{
		Total: count,
		List:  list,
	}, nil
}
