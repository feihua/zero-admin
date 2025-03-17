package company_address

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressReceiveStatusLogic 更新公司默认收货地址状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:25
*/
type UpdateCompanyAddressReceiveStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompanyAddressReceiveStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressReceiveStatusLogic {
	return &UpdateCompanyAddressReceiveStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCompanyAddressReceiveStatus 更新公司默认收货地址状态
func (l *UpdateCompanyAddressReceiveStatusLogic) UpdateCompanyAddressReceiveStatus(req *types.UpdateCompanyAddressReceiveStatusReq) (resp *types.UpdateCompanyAddressStatusResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.UpdateCompanyAddressReceiveStatus(l.ctx, &omsclient.UpdateCompanyAddressReceiveStatusReq{
		Id:            req.Id,
		ReceiveStatus: req.ReceiveStatus, // 是否默认收货地址：0->否；1->是
		UpdateBy:      l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新公司默认收货地址状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateCompanyAddressStatusResp{
		Code:    "000000",
		Message: "更新公司默认收货地址状态成功",
	}, nil
}
