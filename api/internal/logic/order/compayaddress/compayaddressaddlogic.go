package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressAddLogic {
	return CompayAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressAddLogic) CompayAddressAdd(req types.AddCompayAddressReq) (*types.AddCompayAddressResp, error) {
	_, err := l.svcCtx.Oms.CompanyAddressAdd(l.ctx, &omsclient.CompanyAddressAddReq{
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
		return nil, errorx.NewDefaultError("添加公司收发货地址失败")
	}

	return &types.AddCompayAddressResp{
		Code:    "000000",
		Message: "添加公司收发货地址成功",
	}, nil
}
