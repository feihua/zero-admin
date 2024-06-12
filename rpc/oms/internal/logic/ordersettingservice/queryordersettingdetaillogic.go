package ordersettingservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderSettingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingDetailLogic {
	return &QueryOrderSettingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询订单设置表详情
func (l *QueryOrderSettingDetailLogic) QueryOrderSettingDetail(in *omsclient.QueryOrderSettingDetailReq) (*omsclient.QueryOrderSettingDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryOrderSettingDetailResp{}, nil
}
