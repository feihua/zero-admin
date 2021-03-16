package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Oms.CompanyAddressUpdate(l.ctx, &omsclient.CompanyAddressUpdateReq{
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
		return nil, err
	}

	return &types.UpdateCompayAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
