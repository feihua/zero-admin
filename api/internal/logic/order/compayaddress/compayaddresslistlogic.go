package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompayAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressListLogic {
	return CompayAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressListLogic) CompayAddressList(req types.ListCompayAddressReq) (*types.ListCompayAddressResp, error) {
	resp, err := l.svcCtx.Oms.CompanyAddressList(l.ctx, &omsclient.CompanyAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询公司收发货地址列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询公司收发货地址失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtCompayAddressData

	for _, item := range resp.List {
		list = append(list, &types.ListtCompayAddressData{
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
		})
	}

	return &types.ListCompayAddressResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询公司收发货地址成功",
	}, nil
}
