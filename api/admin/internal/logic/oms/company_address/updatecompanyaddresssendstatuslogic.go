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

// UpdateCompanyAddressSendStatusLogic 更新公司默认发货地址状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:29
*/
type UpdateCompanyAddressSendStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompanyAddressSendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressSendStatusLogic {
	return &UpdateCompanyAddressSendStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCompanyAddressSendStatus 更新公司默认发货地址状态
func (l *UpdateCompanyAddressSendStatusLogic) UpdateCompanyAddressSendStatus(req *types.UpdateCompanyAddressSendStatusReq) (resp *types.UpdateCompanyAddressStatusResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.UpdateCompanyAddressSendStatus(l.ctx, &omsclient.UpdateCompanyAddressSendStatusReq{
		Id:         req.Id,
		SendStatus: req.SendStatus, // 默认发货地址：0->否；1->是
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新公司默认发货地址状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateCompanyAddressStatusResp{
		Code:    "000000",
		Message: "更新公司默认发货地址状态成功",
	}, nil
}
