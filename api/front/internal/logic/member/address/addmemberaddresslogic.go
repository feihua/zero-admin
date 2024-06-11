package address

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberAddressLogic 添加收货地址
/*
Author: LiuFeiHua
Date: 2024/5/16 13:58
*/
type AddMemberAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberAddressLogic {
	return &AddMemberAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberAddress 添加收货地址
func (l *AddMemberAddressLogic) AddMemberAddress(req *types.AddMemberAddressReq) (resp *types.AddMemberAddressResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.MemberReceiveAddressService.AddMemberReceiveAddress(l.ctx, &umsclient.AddMemberReceiveAddressReq{
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

	return &types.AddMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
