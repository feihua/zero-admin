package companyaddressservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCompanyAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCompanyAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressDetailLogic {
	return &QueryCompanyAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询公司收发货地址表详情
func (l *QueryCompanyAddressDetailLogic) QueryCompanyAddressDetail(in *omsclient.QueryCompanyAddressDetailReq) (*omsclient.QueryCompanyAddressDetailResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.QueryCompanyAddressDetailResp{}, nil
}
