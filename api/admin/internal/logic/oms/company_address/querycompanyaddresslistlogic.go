package company_address

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCompanyAddressListLogic 查询公司收发货地址列表
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:43
*/
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

// QueryCompanyAddressList 查询公司收发货地址列表
func (l *QueryCompanyAddressListLogic) QueryCompanyAddressList(req *types.QueryCompanyAddressListReq) (resp *types.QueryCompanyAddressListResp, err error) {
	result, err := l.svcCtx.CompanyAddressService.QueryCompanyAddressList(l.ctx, &omsclient.QueryCompanyAddressListReq{
		PageNum:       req.Current,
		PageSize:      req.PageSize,
		AddressName:   req.AddressName,   // 地址名称
		Name:          req.Name,          // 收发货人姓名
		Phone:         req.Phone,         // 收货人电话
		SendStatus:    req.SendStatus,    // 默认发货地址：0->否；1->是
		ReceiveStatus: req.ReceiveStatus, // 默认收货地址：0->否；1->是
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字公司收发货地址列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryCompanyAddressListData

	for _, detail := range result.List {
		list = append(list, &types.QueryCompanyAddressListData{
			Id:            detail.Id,            // 主键ID
			AddressName:   detail.AddressName,   // 地址名称
			Name:          detail.Name,          // 收发货人姓名
			Phone:         detail.Phone,         // 收货人电话
			Province:      detail.Province,      // 省/直辖市
			City:          detail.City,          // 市
			Region:        detail.Region,        // 区
			DetailAddress: detail.DetailAddress, // 详细地址
			SendStatus:    detail.SendStatus,    // 默认发货地址：0->否；1->是
			ReceiveStatus: detail.ReceiveStatus, // 默认收货地址：0->否；1->是
			CreateBy:      detail.CreateBy,      // 创建人ID
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateBy:      detail.UpdateBy,      // 更新人ID
			UpdateTime:    detail.UpdateTime,    // 更新时间

		})
	}

	return &types.QueryCompanyAddressListResp{
		Code:     "000000",
		Message:  "查询公司收发货地址列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
