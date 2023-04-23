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

type AddressDefaultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressDefaultLogic {
	return &AddressDefaultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressDefaultLogic) AddressDefault(req *types.AdressDefaultReq) (resp *types.AdressDefaultResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	addressQueryDetail, err := l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		MemberId:  memberId,
		AddressID: req.AddressID,
	})

	if addressQueryDetail == nil || err != nil {
		// reqStr, _ := json.Marshal(req)
		// logx.WithContext(l.ctx).Errorf("查询会员收货详细地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AdressDefaultResp{
			Code: 1,
			Msg:  "参数有误",
		}, nil
	}

	//
	_, err = l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.AddressID,
		DefaultStatus: 1,
		MemberId:      memberId,
		Name:          addressQueryDetail.Name,
		PhoneNumber:   addressQueryDetail.PhoneNumber,
		PostCode:      addressQueryDetail.PostCode,
		Province:      addressQueryDetail.Province,
		City:          addressQueryDetail.City,
		Region:        addressQueryDetail.Region,
		DetailAddress: addressQueryDetail.DetailAddress,
		Address:       addressQueryDetail.Address,
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员收货默认地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AdressDefaultResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}

	return &types.AdressDefaultResp{
		Code: 0,
		Msg:  "更新会员收货地址成功",
	}, nil
	return
}
