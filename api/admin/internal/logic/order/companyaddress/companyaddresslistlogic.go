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

// CompanyAddressListLogic 公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/13 10:33
*/
type CompanyAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompanyAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressListLogic {
	return &CompanyAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CompanyAddressList 查询公司收货地址
func (l *CompanyAddressListLogic) CompanyAddressList(req *types.ListCompanyAddressReq) (resp *types.ListCompanyAddressResp, err error) {
	result, err := l.svcCtx.CompanyAddressService.QueryCompanyAddressList(l.ctx, &omsclient.QueryCompanyAddressListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询公司收发货地址列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询公司收发货地址失败")
	}

	var list []*types.ListCompanyAddressData

	for _, item := range result.List {
		list = append(list, &types.ListCompanyAddressData{
			Id:            item.Id,
			AddressName:   item.AddressName,
			SendStatus:    item.SendStatus,
			ReceiveStatus: item.ReceiveStatus,
			Name:          item.Name,
			Phone:         item.Phone,
			Province:      item.Province,
			City:          item.City,
			Region:        item.Region,
			DetailAddress: item.DetailAddress,
			CreateBy:      item.CreateBy,
			CreateTime:    item.CreateTime,
			UpdateBy:      item.UpdateBy,
			UpdateTime:    item.UpdateTime,
		})
	}

	return &types.ListCompanyAddressResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询公司收发货地址成功",
	}, nil
}
