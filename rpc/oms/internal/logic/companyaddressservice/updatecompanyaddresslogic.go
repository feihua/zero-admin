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
		ID:            in.Id,            //
		AddressName:   in.AddressName,   // 地址名称
		SendStatus:    in.SendStatus,    // 默认发货地址：0->否；1->是
		ReceiveStatus: in.ReceiveStatus, // 是否默认收货地址：0->否；1->是
		Name:          in.Name,          // 收发货人姓名
		Phone:         in.Phone,         // 收货人电话
		Province:      in.Province,      // 省/直辖市
		City:          in.City,          // 市
		Region:        in.Region,        // 区
		DetailAddress: in.DetailAddress, // 详细地址
		UpdateBy:      in.UpdateBy,      // 更新者
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressResp{}, nil
}
