package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddressListLogic {
	return AddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressListLogic) AddressList(req types.AddressListReq) (resp *types.AddressListResp, err error) {

	reslut, _ := l.svcCtx.Ums.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserId:   req.UserId,
	})

	var list []types.AddressList

	for _, address := range reslut.List {
		list = append(list, types.AddressList{
			ID:            address.Id,
			Name:          address.Name,
			UserID:        address.MemberId,
			Province:      address.Province,
			City:          address.City,
			County:        "",
			AddressDetail: address.DetailAddress,
			AreaCode:      "",
			Tel:           address.PhoneNumber,
			IsDefault:     false,
		})
	}

	return &types.AddressListResp{
		Errno: 0,
		Data: types.AddressListData{
			Total: 0,
			Pages: 0,
			Limit: 0,
			Page:  0,
			List:  list,
		},
		Errmsg: "",
	}, nil
}
