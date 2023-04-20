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

func (l *AddressSaveLogic) AddressSave(req *types.AddressSaveReq) (resp *types.AddressSaveResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.Ums.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
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
		logx.WithContext(l.ctx).Errorf("新增会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressSaveResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}

	return &types.AddressSaveResp{
		Code: 0,
		Msg:  "新增会员收货地址成功",
	}, nil
}
