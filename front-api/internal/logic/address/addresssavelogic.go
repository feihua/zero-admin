package address

import (
	"context"
	"encoding/json"
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
		//新增会员地址
		return addMemberAddress(req, l)
	} else {
		//更新会员地址
		return updateMemberAddress(req, l)
	}
}

//更新会员地址
func updateMemberAddress(req types.AddressSaveReq, l *AddressSaveLogic) (resp *types.AddressSaveResp, err error) {
	_, err = l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.ID,
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

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressSaveResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.AddressSaveResp{
		Errno:  0,
		Errmsg: "更新会员收货地址成功",
	}, nil
}

//新增会员地址
func addMemberAddress(req types.AddressSaveReq, l *AddressSaveLogic) (resp *types.AddressSaveResp, err error) {
	_, err = l.svcCtx.Ums.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
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

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("新增会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressSaveResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.AddressSaveResp{
		Errno:  0,
		Errmsg: "新增会员收货地址成功",
	}, nil
}
