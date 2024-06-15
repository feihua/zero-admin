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

// AddCompanyAddressLogic 添加公司收货地址
/*
Author: LiuFeiHua
Date: 2024/6/15 11:22
*/
type AddCompanyAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompanyAddressLogic {
	return &AddCompanyAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddCompanyAddress 添加公司收货地址
func (l *AddCompanyAddressLogic) AddCompanyAddress(req *types.AddCompanyAddressReq) (resp *types.AddCompanyAddressResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.AddCompanyAddress(l.ctx, &omsclient.AddCompanyAddressReq{
		AddressName:   req.AddressName,
		SendStatus:    req.SendStatus,
		ReceiveStatus: req.ReceiveStatus,
		Name:          req.Name,
		Phone:         req.Phone,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
		CreateBy:      l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加公司收发货地址信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加公司收发货地址失败")
	}

	return &types.AddCompanyAddressResp{
		Code:    "000000",
		Message: "添加公司收发货地址成功",
	}, nil
}
