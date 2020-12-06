package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompanyAddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressUpdateLogic {
	return &CompanyAddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressUpdateLogic) CompanyAddressUpdate(in *oms.CompanyAddressUpdateReq) (*oms.CompanyAddressUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &oms.CompanyAddressUpdateResp{}, nil
}
