package companyaddressservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompanyAddressStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressStatusLogic {
	return &UpdateCompanyAddressStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新公司收发货地址表状态
func (l *UpdateCompanyAddressStatusLogic) UpdateCompanyAddressStatus(in *omsclient.UpdateCompanyAddressStatusReq) (*omsclient.UpdateCompanyAddressStatusResp, error) {
	// todo: add your logic here and delete this line

	return &omsclient.UpdateCompanyAddressStatusResp{}, nil
}
