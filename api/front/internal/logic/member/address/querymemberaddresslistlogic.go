package address

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberAddressListLogic 获取所有收货地址
/*
Author: LiuFeiHua
Date: 2024/5/16 13:58
*/
type QueryMemberAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberAddressListLogic {
	return &QueryMemberAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberAddressList 获取所有收货地址
func (l *QueryMemberAddressListLogic) QueryMemberAddressList() (resp *types.ListMemberAddressResp, err error) {
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
			Name:          member.MemberName,
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
