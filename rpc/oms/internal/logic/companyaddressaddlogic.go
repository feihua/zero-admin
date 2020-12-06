package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompanyAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressAddLogic {
	return &CompanyAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressAddLogic) CompanyAddressAdd(in *oms.CompanyAddressAddReq) (*oms.CompanyAddressAddResp, error) {
	// todo: add your logic here and delete this line

	return &oms.CompanyAddressAddResp{}, nil
}
