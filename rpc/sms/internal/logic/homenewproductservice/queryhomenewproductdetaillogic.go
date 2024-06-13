package homenewproductservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHomeNewProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeNewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeNewProductDetailLogic {
	return &QueryHomeNewProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询新鲜好物表详情
func (l *QueryHomeNewProductDetailLogic) QueryHomeNewProductDetail(in *smsclient.QueryHomeNewProductDetailReq) (*smsclient.QueryHomeNewProductDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryHomeNewProductDetailResp{}, nil
}
