package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CompanyAddressAddLogic 公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/13 10:39
*/
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

// CompanyAddressAdd 添加公司收货地址
func (l *CompanyAddressAddLogic) CompanyAddressAdd(in *omsclient.CompanyAddressAddReq) (*omsclient.CompanyAddressAddResp, error) {
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

	return &omsclient.CompanyAddressAddResp{}, nil
}
