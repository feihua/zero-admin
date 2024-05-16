package address

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
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
		return nil, err
	}

	return &types.UpdateMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
