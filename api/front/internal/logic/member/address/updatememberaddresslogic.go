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

// UpdateMemberAddressLogic 修改收货地址
/*
Author: LiuFeiHua
Date: 2024/5/16 13:59
*/
type UpdateMemberAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberAddressLogic {
	return &UpdateMemberAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberAddress 修改收货地址
func (l *UpdateMemberAddressLogic) UpdateMemberAddress(req *types.UpdateMemberAddressReq) (resp *types.UpdateMemberAddressResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberReceiveAddressService.UpdateMemberReceiveAddress(l.ctx, &umsclient.UpdateMemberReceiveAddressReq{
		Id:            req.Id,
		MemberId:      memberId,
		MemberName:    req.Name,
		PhoneNumber:   req.PhoneNumber,
		DefaultStatus: req.DefaultStatus,
		PostCode:      req.PostCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})
	if err != nil {
		logc.Errorf(l.ctx, "修改收货地址失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
