package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddressSaveLogic {
	return AddressSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressSaveLogic) AddressSave(req types.AddressSaveReq) (resp *types.AddressSaveResp, err error) {

	if req.ID == 0 {
		l.svcCtx.Ums.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
			MemberId:      req.UserID,
			Name:          req.Name,
			PhoneNumber:   req.Tel,
			DefaultStatus: 0,
			PostCode:      req.PostalCode,
			Province:      req.Province,
			City:          req.City,
			Region:        req.AreaCode,
			DetailAddress: req.AddressDetail,
		})
	} else {
		l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
			MemberId:      req.UserID,
			Name:          req.Name,
			PhoneNumber:   req.Tel,
			DefaultStatus: 0,
			PostCode:      req.PostalCode,
			Province:      req.Province,
			City:          req.City,
			Region:        req.AreaCode,
			DetailAddress: req.AddressDetail,
		})
	}

	return &types.AddressSaveResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
