package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddressDetailLogic {
	return AddressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressDetailLogic) AddressDetail(req types.AddressDetailReq) (resp *types.AddressDetailResp, err error) {

	addressQueryDetail, _ := l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		UserId:    req.UserID,
		AddressID: req.AddressID,
	})

	data := types.AddressDetailData{
		IsDeleted:     false,
		NotDeleted:    false,
		ID:            addressQueryDetail.Id,
		Name:          addressQueryDetail.Name,
		UserID:        addressQueryDetail.MemberId,
		Province:      addressQueryDetail.Province,
		City:          addressQueryDetail.City,
		County:        "",
		AddressDetail: addressQueryDetail.DetailAddress,
		AreaCode:      addressQueryDetail.Region,
		PostalCode:    addressQueryDetail.PostCode,
		Tel:           addressQueryDetail.PhoneNumber,
		IsDefault:     true,
		//AddTime:       addressQueryDetail,
		//UpdateTime:    "",
		//Deleted:       false,
	}

	return &types.AddressDetailResp{
		Errno:  0,
		Data:   data,
		Errmsg: "",
	}, nil
}
