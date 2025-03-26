package companyaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCompanyAddressListLogic 查询公司收发货地址列表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:11
*/
type QueryCompanyAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCompanyAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressListLogic {
	return &QueryCompanyAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCompanyAddressList 查询公司收发货地址列表
func (l *QueryCompanyAddressListLogic) QueryCompanyAddressList(in *omsclient.QueryCompanyAddressListReq) (*omsclient.QueryCompanyAddressListResp, error) {
	q := query.OmsCompanyAddress.WithContext(l.ctx)

	if len(in.AddressName) > 0 {
		q = q.Where(query.OmsCompanyAddress.AddressName.Like("%" + in.AddressName + "%"))
	}
	if len(in.Name) > 0 {
		q = q.Where(query.OmsCompanyAddress.Name.Like("%" + in.Name + "%"))
	}
	if len(in.Phone) > 0 {
		q = q.Where(query.OmsCompanyAddress.Phone.Like("%" + in.Phone + "%"))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询公司收发货地址列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询公司收发货地址列表失败")
	}

	var list []*omsclient.CompanyAddressListData
	for _, item := range result {

		list = append(list, &omsclient.CompanyAddressListData{
			Id:            item.ID,                                 //
			AddressName:   item.AddressName,                        // 地址名称
			SendStatus:    item.SendStatus,                         // 默认发货地址：0->否；1->是
			ReceiveStatus: item.ReceiveStatus,                      // 是否默认收货地址：0->否；1->是
			Name:          item.Name,                               // 收发货人姓名
			Phone:         item.Phone,                              // 收货人电话
			Province:      item.Province,                           // 省/直辖市
			City:          item.City,                               // 市
			Region:        item.Region,                             // 区
			DetailAddress: item.DetailAddress,                      // 详细地址
			CreateBy:      item.CreateBy,                           // 创建者
			CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:      item.UpdateBy,                           // 更新者
			UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &omsclient.QueryCompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil

}
