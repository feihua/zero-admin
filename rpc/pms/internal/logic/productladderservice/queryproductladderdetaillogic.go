package productladderservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductLadderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductLadderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLadderDetailLogic {
	return &QueryProductLadderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询产品阶梯价格表(只针对同商品)详情
func (l *QueryProductLadderDetailLogic) QueryProductLadderDetail(in *pmsclient.QueryProductLadderDetailReq) (*pmsclient.QueryProductLadderDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryProductLadderDetailResp{}, nil
}
