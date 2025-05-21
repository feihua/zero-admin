package address

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberAddressStatusLogic 更新会员收货地址状态状态
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type UpdateMemberAddressStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberAddressStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberAddressStatusLogic {
	return &UpdateMemberAddressStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberAddressStatus 更新会员收货地址状态
func (l *UpdateMemberAddressStatusLogic) UpdateMemberAddressStatus(req *types.UpdateMemberAddressStatusReq) (resp *types.UpdateMemberAddressStatusResp, err error) {
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

	return &types.UpdateMemberAddressStatusResp{
		Code:    0,
		Message: "更新会员收货地址状态成功",
	}, nil
}
