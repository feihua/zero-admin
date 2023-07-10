package companyaddressservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompanyAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanyAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanyAddressListLogic {
	return &CompanyAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompanyAddressListLogic) CompanyAddressList(in *omsclient.CompanyAddressListReq) (*omsclient.CompanyAddressListResp, error) {
	all, err := l.svcCtx.OmsCompanyAddressModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsCompanyAddressModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询公司收发货地址列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*omsclient.CompanyAddressListData
	for _, item := range *all {

		list = append(list, &omsclient.CompanyAddressListData{
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询公司收发货地址列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &omsclient.CompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil
}
