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

// QueryCompanyAddressDetailLogic 查询公司收发货地址表详情
/*
Author: 刘飞华
Date: 2025/02/05 11:40:41
*/
type QueryCompanyAddressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCompanyAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressDetailLogic {
	return &QueryCompanyAddressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCompanyAddressDetail 查询公司收发货地址表详情
func (l *QueryCompanyAddressDetailLogic) QueryCompanyAddressDetail(req *types.QueryCompanyAddressDetailReq) (resp *types.QueryCompanyAddressDetailResp, err error) {

	detail, err := l.svcCtx.CompanyAddressService.QueryCompanyAddressDetail(l.ctx, &omsclient.QueryCompanyAddressDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询公司收发货地址表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryCompanyAddressDetailData{
		Id:            detail.Id,            //
		AddressName:   detail.AddressName,   // 地址名称
		SendStatus:    detail.SendStatus,    // 默认发货地址：0->否；1->是
		ReceiveStatus: detail.ReceiveStatus, // 是否默认收货地址：0->否；1->是
		Name:          detail.Name,          // 收发货人姓名
		Phone:         detail.Phone,         // 收货人电话
		Province:      detail.Province,      // 省/直辖市
		City:          detail.City,          // 市
		Region:        detail.Region,        // 区
		DetailAddress: detail.DetailAddress, // 详细地址
		CreateBy:      detail.CreateBy,      // 创建者
		CreateTime:    detail.CreateTime,    // 创建时间
		UpdateBy:      detail.UpdateBy,      // 更新者
		UpdateTime:    detail.UpdateTime,    // 更新时间
	}
	return &types.QueryCompanyAddressDetailResp{
		Code:    "000000",
		Message: "查询公司收发货地址表成功",
		Data:    data,
	}, nil
}
