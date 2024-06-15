package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressSendStatusLogic 更新公司收发货地址表状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:35
*/
type UpdateCompanyAddressSendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressSendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressSendStatusLogic {
	return &UpdateCompanyAddressSendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCompanyAddressSendStatus 更新公司收发货地址表状态
func (l *UpdateCompanyAddressSendStatusLogic) UpdateCompanyAddressSendStatus(in *omsclient.UpdateCompanyAddressSendStatusReq) (*omsclient.UpdateCompanyAddressStatusResp, error) {
	q := query.OmsCompanyAddress
	if in.SendStatus == 1 {
		_, err := q.WithContext(l.ctx).Where(q.SendStatus.Eq(1)).Update(q.SendStatus, 0)
		if err != nil {
			return nil, err
		}
	}
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.SendStatus, in.SendStatus)

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressStatusResp{}, nil
}
