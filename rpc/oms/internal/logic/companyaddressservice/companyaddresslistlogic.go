package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.OmsCompanyAddress.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询公司收发货地址列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.CompanyAddressListData
	for _, item := range result {

		list = append(list, &omsclient.CompanyAddressListData{
			Id:            item.ID,
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

	logc.Infof(l.ctx, "查询公司收发货地址列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.CompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil
}
