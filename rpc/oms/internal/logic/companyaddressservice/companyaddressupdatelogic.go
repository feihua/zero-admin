package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CompanyAddressUpdateLogic 公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/13 10:41
*/
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

// CompanyAddressUpdate 更新删除公司收货地址
func (l *CompanyAddressUpdateLogic) CompanyAddressUpdate(in *omsclient.CompanyAddressUpdateReq) (*omsclient.CompanyAddressUpdateResp, error) {
	q := query.OmsCompanyAddress
	_, err := q.WithContext(l.ctx).Updates(&model.OmsCompanyAddress{
		ID:            in.Id,
		AddressName:   in.AddressName,
		SendStatus:    in.SendStatus,
		ReceiveStatus: in.ReceiveStatus,
		Name:          in.Name,
		Phone:         in.Phone,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
		UpdateBy:      &in.UpdateBy,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.CompanyAddressUpdateResp{}, nil
}
