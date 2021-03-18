package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *CompanyAddressListLogic) CompanyAddressList(in *oms.CompanyAddressListReq) (*oms.CompanyAddressListResp, error) {
	all, _ := l.svcCtx.OmsCompanyAddressModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsCompanyAddressModel.Count()

	var list []*oms.CompanyAddressListData
	for _, item := range *all {

		list = append(list, &oms.CompanyAddressListData{
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

	fmt.Println(list)
	return &oms.CompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil
}
