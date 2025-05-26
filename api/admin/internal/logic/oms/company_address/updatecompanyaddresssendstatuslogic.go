package company_address

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
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
func (l *UpdateCompanyAddressSendStatusLogic) UpdateCompanyAddressSendStatus(req *types.UpdateCompanyAddressStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CompanyAddressService.UpdateCompanyAddressSendStatus(l.ctx, &omsclient.UpdateCompanyAddressStatusReq{
		Id:       req.Id,
		Status:   req.Status, // 默认发货地址：0->否；1->是
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新公司默认发货地址状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
