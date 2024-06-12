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

// CompanyAddressDeleteLogic 公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/13 10:32
*/
type CompanyAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressDeleteLogic {
	return &CompanyAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CompanyAddressDelete 删除公司收货地址
func (l *CompanyAddressDeleteLogic) CompanyAddressDelete(req *types.DeleteCompanyAddressReq) (resp *types.DeleteCompanyAddressResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.DeleteCompanyAddress(l.ctx, &omsclient.DeleteCompanyAddressReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除公司地址异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除公司地址失败")
	}

	return &types.DeleteCompanyAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
