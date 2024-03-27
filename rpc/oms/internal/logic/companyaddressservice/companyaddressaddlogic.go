package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/omsmodel"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CompanyAddressAddLogic) CompanyAddressAdd(in *omsclient.CompanyAddressAddReq) (*omsclient.CompanyAddressAddResp, error) {
	_, err := l.svcCtx.OmsCompanyAddressModel.Insert(l.ctx, &omsmodel.OmsCompanyAddress{
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

	return &omsclient.CompanyAddressAddResp{}, nil
}
