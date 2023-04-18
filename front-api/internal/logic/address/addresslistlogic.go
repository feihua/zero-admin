package address

import (
	"context"
	"encoding/json"
	"zero-admin/common/ctxdata"
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
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	result, err := l.svcCtx.Ums.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		MemberId: memberId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询用户收货地址列表失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressListResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}

	var list []types.AddressList

	for _, address := range result.List {
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
		Code: 0,
		Data: types.AddressListData{
			Total: 0,
			Pages: 0,
			Limit: 0,
			Page:  0,
			List:  list,
		},
		Msg: "查询用户收货地址列表成功",
	}, nil
}
