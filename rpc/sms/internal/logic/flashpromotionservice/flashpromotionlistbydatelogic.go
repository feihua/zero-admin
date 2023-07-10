package flashpromotionservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionListByDateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionListByDateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListByDateLogic {
	return &FlashPromotionListByDateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionListByDateLogic) FlashPromotionListByDate(in *smsclient.FlashPromotionListByDateReq) (*smsclient.FlashPromotionListByDateResp, error) {
	all, err := l.svcCtx.SmsFlashPromotionModel.FindAllByCurrentDate(l.ctx, in.CurrentDate)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionListData
	for _, item := range *all {

		list = append(list, &smsclient.FlashPromotionListData{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02"),
			EndDate:    item.EndDate.Format("2006-01-02"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询限时购列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &smsclient.FlashPromotionListByDateResp{List: list}, nil
}
