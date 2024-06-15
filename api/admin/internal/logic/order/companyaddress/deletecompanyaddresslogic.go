package companyaddress

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCompanyAddressLogic 删除公司收货地址
/*
Author: LiuFeiHua
Date: 2024/6/15 11:22
*/
type DeleteCompanyAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanyAddressLogic {
	return &DeleteCompanyAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCompanyAddress 删除公司收货地址
func (l *DeleteCompanyAddressLogic) DeleteCompanyAddress(req *types.DeleteCompanyAddressReq) (resp *types.DeleteCompanyAddressResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.DeleteCompanyAddress(l.ctx, &omsclient.DeleteCompanyAddressReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除公司地址异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除公司地址失败")
	}

	return &types.DeleteCompanyAddressResp{
		Code:    "000000",
		Message: "删除公司收货地址成功",
	}, nil
}
