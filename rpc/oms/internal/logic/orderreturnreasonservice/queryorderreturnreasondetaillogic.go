package orderreturnreasonservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderReturnReasonDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnReasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonDetailLogic {
	return &QueryOrderReturnReasonDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询退货原因表详情
func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(in *omsclient.QueryOrderReturnReasonDetailReq) (*omsclient.QueryOrderReturnReasonDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryOrderReturnReasonDetailResp{}, nil
}
