package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// IntegrationConsumeSettingListLogic 查询积分消费设置列表
/*
Author: LiuFeiHua
Date: 2024/5/23 13:44
*/
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

// IntegrationConsumeSettingList 查询积分消费设置列表
func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(in *umsclient.IntegrationConsumeSettingListReq) (*umsclient.IntegrationConsumeSettingListResp, error) {
	q := query.UmsIntegrationConsumeSetting.WithContext(l.ctx)
	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.IntegrationConsumeSettingListData
	for _, item := range result {

		list = append(list, &umsclient.IntegrationConsumeSettingListData{
			Id:                 item.ID,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	logc.Infof(l.ctx, "查询积分消费设置列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.IntegrationConsumeSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
