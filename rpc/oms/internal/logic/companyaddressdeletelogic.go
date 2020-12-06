package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &oms.CompanyAddressDeleteResp{}, nil
}
