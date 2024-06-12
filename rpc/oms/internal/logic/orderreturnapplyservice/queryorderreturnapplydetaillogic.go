package orderreturnapplyservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderReturnApplyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnApplyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyDetailLogic {
	return &QueryOrderReturnApplyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询订单退货申请详情
func (l *QueryOrderReturnApplyDetailLogic) QueryOrderReturnApplyDetail(in *omsclient.QueryOrderReturnApplyDetailReq) (*omsclient.QueryOrderReturnApplyDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryOrderReturnApplyDetailResp{}, nil
}
