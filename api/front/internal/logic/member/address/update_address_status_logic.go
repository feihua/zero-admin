package address

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateAddressStatusLogic 更新会员默认收货地址
/*
Author: LiuFeiHua
Date: 2025/6/19 10:54
*/
func NewUpdateAddressStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressStatusLogic {
	return &UpdateAddressStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateAddressStatus 更新会员默认收货地址
func (l *UpdateAddressStatusLogic) UpdateAddressStatus(req *types.UpdateAddressStatusReq) (resp *types.AddressResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberAddressService.UpdateMemberAddressStatus(l.ctx, &umsclient.UpdateMemberAddressStatusReq{
		Id:        req.Id, // 主键ID
		IsDefault: req.IsDefault,
		MemberId:  memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员收货地址状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddressResp{
		Code:    0,
		Message: "更新会员收货地址状态成功",
	}, nil
}
