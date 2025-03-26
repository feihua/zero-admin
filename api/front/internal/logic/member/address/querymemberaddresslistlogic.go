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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	addressList, err := l.svcCtx.MemberReceiveAddressService.QueryMemberReceiveAddressList(l.ctx, &umsclient.QueryMemberReceiveAddressListReq{
		PageNum:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "获取所有收货地址失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

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
