package address

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberAddressListLogic 收货地址相关
/*
Author: LiuFeiHua
Date: 2023/11/29 16:13
*/
type MemberAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressListLogic {
	return &MemberAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MemberAddressList 查询会员收货地址
func (l *MemberAddressListLogic) MemberAddressList() (resp *types.ListMemberAddressResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	addressList, _ := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	var list []types.ListMemberAddressData

	for _, member := range addressList.List {
		list = append(list, types.ListMemberAddressData{
			Id:            member.Id,
			MemberId:      member.MemberId,
			Name:          member.Name,
			PhoneNumber:   member.PhoneNumber,
			DefaultStatus: member.DefaultStatus,
			PostCode:      member.PostCode,
			Province:      member.Province,
			City:          member.City,
			Region:        member.Region,
			DetailAddress: member.DetailAddress,
		})
	}

	return &types.ListMemberAddressResp{
		Data:    list,
		Success: true,
		Code:    0,
		Message: "操作成功",
	}, nil
}
