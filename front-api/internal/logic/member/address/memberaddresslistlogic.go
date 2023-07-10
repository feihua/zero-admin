package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberAddressListLogic) MemberAddressList(req *types.ListMemberAddressReq) (resp *types.ListMemberAddressResp, err error) {
	addressList, _ := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		MemberId: l.ctx.Value("memberId").(int64),
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
		Current:  0,
		Data:     list,
		PageSize: 0,
		Success:  false,
		Total:    0,
		Code:     0,
		Message:  "",
	}, nil
}
