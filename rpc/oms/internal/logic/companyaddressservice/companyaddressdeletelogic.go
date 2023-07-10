package companyaddressservicelogic

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressDeleteLogic {
	return &CompanyAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressDeleteLogic) CompanyAddressDelete(in *omsclient.CompanyAddressDeleteReq) (*omsclient.CompanyAddressDeleteResp, error) {
	err := l.svcCtx.OmsCompanyAddressModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &omsclient.CompanyAddressDeleteResp{}, nil
}
