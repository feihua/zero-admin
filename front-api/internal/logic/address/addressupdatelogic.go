package address

import (
	"context"
	"encoding/json"

	"zero-admin/common/ctxdata"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressUpdateLogic {
	return &AddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressUpdateLogic) AddressUpdate(req *types.AddressUpdateReq) (resp *types.AddressUpdateResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)

	addressQueryDetail, err := l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		MemberId:  memberId,
		AddressID: req.AddressID,
	})

	if addressQueryDetail == nil || err != nil {
		// reqStr, _ := json.Marshal(req)
		// logx.WithContext(l.ctx).Errorf("查询会员收货详细地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressUpdateResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}
	_, err = l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.AddressID,
		MemberId:      memberId,
		Name:          req.Name,
		PhoneNumber:   req.Tel,
		DefaultStatus: 0,
		PostCode:      req.PostalCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.AddressDetail,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressUpdateResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}

	return &types.AddressUpdateResp{
		Code: 0,
		Msg:  "更新会员收货地址成功",
	}, nil
}
