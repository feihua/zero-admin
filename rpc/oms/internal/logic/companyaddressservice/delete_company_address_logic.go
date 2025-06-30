package companyaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCompanyAddressLogic 删除公司收发货地址
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:43
*/
type DeleteCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanyAddressLogic {
	return &DeleteCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCompanyAddress 删除公司收发货地址
func (l *DeleteCompanyAddressLogic) DeleteCompanyAddress(in *omsclient.DeleteCompanyAddressReq) (*omsclient.DeleteCompanyAddressResp, error) {
	q := query.OmsCompanyAddress

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除公司收发货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除公司收发货地址失败")
	}

	return &omsclient.DeleteCompanyAddressResp{}, nil
}
