package cartitemservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCartItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCartItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCartItemDetailLogic {
	return &QueryCartItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询购物车详情
func (l *QueryCartItemDetailLogic) QueryCartItemDetail(in *omsclient.QueryCartItemDetailReq) (*omsclient.QueryCartItemDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryCartItemDetailResp{}, nil
}
