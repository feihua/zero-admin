package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新公司收发货地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新公司收发货地址失败")
	}

	return &types.UpdateCompayAddressResp{
		Code:    "000000",
		Message: "更新公司收发货地址成功",
	}, nil
}
