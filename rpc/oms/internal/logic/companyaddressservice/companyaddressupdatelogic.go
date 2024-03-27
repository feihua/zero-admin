package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CompanyAddressUpdateLogic) CompanyAddressUpdate(in *omsclient.CompanyAddressUpdateReq) (*omsclient.CompanyAddressUpdateResp, error) {
	err := l.svcCtx.OmsCompanyAddressModel.Update(l.ctx, &omsmodel.OmsCompanyAddress{
		Id:            in.Id,
		AddressName:   in.AddressName,
		SendStatus:    in.SendStatus,
		ReceiveStatus: in.ReceiveStatus,
		Name:          in.Name,
		Phone:         in.Phone,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.CompanyAddressUpdateResp{}, nil
}
