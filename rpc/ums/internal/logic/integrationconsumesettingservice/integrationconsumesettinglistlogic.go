package integrationconsumesettingservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingListLogic {
	return &IntegrationConsumeSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(in *umsclient.IntegrationConsumeSettingListReq) (*umsclient.IntegrationConsumeSettingListResp, error) {
	all, err := l.svcCtx.UmsIntegrationConsumeSettingModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsIntegrationConsumeSettingModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分消费设置列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.IntegrationConsumeSettingListData
	for _, item := range *all {

		list = append(list, &umsclient.IntegrationConsumeSettingListData{
			Id:                 item.Id,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询积分消费设置列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.IntegrationConsumeSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
