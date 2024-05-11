package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompayAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressUpdateLogic {
	return CompayAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressUpdateLogic) CompayAddressUpdate(req types.UpdateCompayAddressReq) (*types.UpdateCompayAddressResp, error) {
	_, err := l.svcCtx.CompanyAddressService.CompanyAddressUpdate(l.ctx, &omsclient.CompanyAddressUpdateReq{
		Id:            req.Id,
		AddressName:   req.AddressName,
		SendStatus:    req.SendStatus,
		ReceiveStatus: req.ReceiveStatus,
		Name:          req.Name,
		Phone:         req.Phone,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新公司收发货地址信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新公司收发货地址失败")
	}

	return &types.UpdateCompayAddressResp{
		Code:    "000000",
		Message: "更新公司收发货地址成功",
	}, nil
}
