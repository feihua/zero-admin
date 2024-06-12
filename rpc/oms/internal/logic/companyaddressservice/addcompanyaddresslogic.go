package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCompanyAddressLogic 添加公司收发货地址表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:10
*/
type AddCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompanyAddressLogic {
	return &AddCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCompanyAddress 添加公司收发货地址表
func (l *AddCompanyAddressLogic) AddCompanyAddress(in *omsclient.AddCompanyAddressReq) (*omsclient.AddCompanyAddressResp, error) {
	err := query.OmsCompanyAddress.WithContext(l.ctx).Create(&model.OmsCompanyAddress{
		AddressName:   in.AddressName,
		SendStatus:    in.SendStatus,
		ReceiveStatus: in.ReceiveStatus,
		Name:          in.Name,
		Phone:         in.Phone,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
		CreateBy:      in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.AddCompanyAddressResp{}, nil
}
