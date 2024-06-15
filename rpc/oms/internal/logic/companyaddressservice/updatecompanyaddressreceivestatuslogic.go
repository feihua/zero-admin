package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressReceiveStatusLogic 更新公司收发货地址表状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:34
*/
type UpdateCompanyAddressReceiveStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressReceiveStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressReceiveStatusLogic {
	return &UpdateCompanyAddressReceiveStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCompanyAddressReceiveStatus 更新公司收发货地址表状态
func (l *UpdateCompanyAddressReceiveStatusLogic) UpdateCompanyAddressReceiveStatus(in *omsclient.UpdateCompanyAddressReceiveStatusReq) (*omsclient.UpdateCompanyAddressStatusResp, error) {
	q := query.OmsCompanyAddress
	if in.ReceiveStatus == 1 {
		_, err := q.WithContext(l.ctx).Where(q.ReceiveStatus.Eq(1)).Update(q.ReceiveStatus, 0)
		if err != nil {
			return nil, err
		}
	}
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.ReceiveStatus, in.ReceiveStatus)

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressStatusResp{}, nil
}
