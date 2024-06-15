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

// UpdateCompanyAddressLogic 更新公司收货地址
/*
Author: LiuFeiHua
Date: 2024/6/15 11:24
*/
type UpdateCompanyAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressLogic {
	return &UpdateCompanyAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCompanyAddress 更新公司收货地址
func (l *UpdateCompanyAddressLogic) UpdateCompanyAddress(req *types.UpdateCompanyAddressReq) (resp *types.UpdateCompanyAddressResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.UpdateCompanyAddress(l.ctx, &omsclient.UpdateCompanyAddressReq{
		Id:            req.Id,
		AddressName:   req.AddressName,
		Name:          req.Name,
		Phone:         req.Phone,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
		UpdateBy:      l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新公司收发货地址信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新公司收发货地址失败")
	}

	return &types.UpdateCompanyAddressResp{
		Code:    "000000",
		Message: "更新公司收发货地址成功",
	}, nil
}
