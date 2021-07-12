package logic

import (
	"context"
	"encoding/json"
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
		reqStr, _ := json.Marshal(req)
		logx.Errorf("添加公司收发货地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加公司收发货地址失败")
	}

	return &types.AddCompayAddressResp{
		Code:    "000000",
		Message: "添加公司收发货地址成功",
	}, nil
}
