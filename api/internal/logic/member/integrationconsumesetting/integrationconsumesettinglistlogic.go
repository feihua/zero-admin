package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingListLogic {
	return IntegrationConsumeSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(req types.ListIntegrationConsumeSettingReq) (*types.ListIntegrationConsumeSettingResp, error) {
	resp, err := l.svcCtx.Ums.IntegrationConsumeSettingList(l.ctx, &umsclient.IntegrationConsumeSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询积分消费设置列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询积分消费设置失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtIntegrationConsumeSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListtIntegrationConsumeSettingData{
			Id:                 item.Id,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	return &types.ListIntegrationConsumeSettingResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询积分消费设置成功",
	}, nil
}
