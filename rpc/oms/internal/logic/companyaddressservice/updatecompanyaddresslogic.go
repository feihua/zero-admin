package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressLogic 更新公司收发货地址表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:11
*/
type UpdateCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressLogic {
	return &UpdateCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCompanyAddress 更新公司收发货地址表
func (l *UpdateCompanyAddressLogic) UpdateCompanyAddress(in *omsclient.UpdateCompanyAddressReq) (*omsclient.UpdateCompanyAddressResp, error) {
	q := query.OmsCompanyAddress
	_, err := q.WithContext(l.ctx).Updates(&model.OmsCompanyAddress{
		ID:            in.Id,
		AddressName:   in.AddressName,
		Name:          in.Name,
		Phone:         in.Phone,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
		UpdateBy:      in.UpdateBy,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressResp{}, nil
}
