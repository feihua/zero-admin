package companyaddress

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCompanyAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCompanyAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressListLogic {
	return &QueryCompanyAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCompanyAddressListLogic) QueryCompanyAddressList(req *types.QueryCompanyAddressListReq) (resp *types.QueryCompanyAddressListResp, err error) {
	result, err := l.svcCtx.CompanyAddressService.QueryCompanyAddressList(l.ctx, &omsclient.QueryCompanyAddressListReq{
		PageNum:     req.Current,
		PageSize:    req.PageSize,
		AddressName: req.AddressName,
		Name:        req.Name,
		Phone:       req.Phone,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询公司收发货地址列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询公司收发货地址失败")
	}

	var list []*types.QueryCompanyAddressListData

	for _, item := range result.List {
		list = append(list, &types.QueryCompanyAddressListData{
			Id:            item.Id,            //
			AddressName:   item.AddressName,   // 地址名称
			SendStatus:    item.SendStatus,    // 默认发货地址：0->否；1->是
			ReceiveStatus: item.ReceiveStatus, // 是否默认收货地址：0->否；1->是
			Name:          item.Name,          // 收发货人姓名
			Phone:         item.Phone,         // 收货人电话
			Province:      item.Province,      // 省/直辖市
			City:          item.City,          // 市
			Region:        item.Region,        // 区
			DetailAddress: item.DetailAddress, // 详细地址
			CreateBy:      item.CreateBy,      // 创建者
			CreateTime:    item.CreateTime,    // 创建时间
			UpdateBy:      item.UpdateBy,      // 更新者
			UpdateTime:    item.UpdateTime,    // 更新时间
		})
	}

	return &types.QueryCompanyAddressListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询公司收发货地址成功",
	}, nil
}
