package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CompanyAddressDeleteLogic 删除公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/8 9:30
*/
type CompanyAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressDeleteLogic {
	return &CompanyAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CompanyAddressDelete 删除公司收货地址
func (l *CompanyAddressDeleteLogic) CompanyAddressDelete(in *omsclient.CompanyAddressDeleteReq) (*omsclient.CompanyAddressDeleteResp, error) {
	q := query.OmsCompanyAddress
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.CompanyAddressDeleteResp{}, nil
}
