package flashpromotionsessionservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

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

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(in *smsclient.FlashPromotionSessionListReq) (*smsclient.FlashPromotionSessionListResp, error) {
	all, err := l.svcCtx.SmsFlashPromotionSessionModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsFlashPromotionSessionModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询限时购场次列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionSessionListData
	for _, item := range *all {

		list = append(list, &smsclient.FlashPromotionSessionListData{
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
	return &smsclient.FlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil
}
