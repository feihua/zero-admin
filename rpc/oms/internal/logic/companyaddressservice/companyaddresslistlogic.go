package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CompanyAddressListLogic 公司收货地址
/*
Author: LiuFeiHua
Date: 2024/5/13 10:42
*/
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

// CompanyAddressList 查询公司收货地址
func (l *CompanyAddressListLogic) CompanyAddressList(in *omsclient.CompanyAddressListReq) (*omsclient.CompanyAddressListResp, error) {
	q := query.OmsCompanyAddress.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询公司收发货地址列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.CompanyAddressListData
	for _, item := range result {

		address := &omsclient.CompanyAddressListData{}
		address.Id = item.ID
		address.AddressName = item.AddressName
		address.SendStatus = item.SendStatus
		address.ReceiveStatus = item.ReceiveStatus
		address.Name = item.Name
		address.Phone = item.Phone
		address.Province = item.Province
		address.City = item.City
		address.Region = item.Region
		address.DetailAddress = item.DetailAddress
		address.CreateBy = item.CreateBy
		address.CreateTime = item.CreateTime.Format("2006-01-02 15:04:05")

		if item.UpdateBy != nil {
			address.UpdateBy = *item.UpdateBy
		}
		if item.UpdateTime != nil {
			address.UpdateTime = item.UpdateTime.Format("2006-01-02 15:04:05")
		}

		list = append(list, address)
	}

	logc.Infof(l.ctx, "查询公司收发货地址列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.CompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil
}
