package logic

import (
	"context"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

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

func (l *CompanyAddressDeleteLogic) CompanyAddressDelete(in *oms.CompanyAddressDeleteReq) (*oms.CompanyAddressDeleteResp, error) {
	err := l.svcCtx.OmsCompanyAddressModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.CompanyAddressDeleteResp{}, nil
}
