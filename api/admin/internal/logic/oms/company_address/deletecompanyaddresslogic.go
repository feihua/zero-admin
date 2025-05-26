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

// DeleteCompanyAddressLogic 删除公司收发货地址
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:43
*/
type DeleteCompanyAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanyAddressLogic {
	return &DeleteCompanyAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCompanyAddress 删除公司收发货地址
func (l *DeleteCompanyAddressLogic) DeleteCompanyAddress(req *types.DeleteCompanyAddressReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CompanyAddressService.DeleteCompanyAddress(l.ctx, &omsclient.DeleteCompanyAddressReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除公司收发货地址失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除公司收发货地址成功",
	}, nil
}
