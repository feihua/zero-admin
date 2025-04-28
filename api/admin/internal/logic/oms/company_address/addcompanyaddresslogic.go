package company_address

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
func (l *AddCompanyAddressLogic) AddCompanyAddress(req *types.AddCompanyAddressReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.AddCompanyAddress(l.ctx, &omsclient.AddCompanyAddressReq{
		AddressName:   req.AddressName,   // 地址名称
		SendStatus:    req.SendStatus,    // 默认发货地址：0->否；1->是
		ReceiveStatus: req.ReceiveStatus, // 是否默认收货地址：0->否；1->是
		Name:          req.Name,          // 收发货人姓名
		Phone:         req.Phone,         // 收货人电话
		Province:      req.Province,      // 省/直辖市
		City:          req.City,          // 市
		Region:        req.Region,        // 区
		DetailAddress: req.DetailAddress, // 详细地址
		CreateBy:      l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加公司收发货地址信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
