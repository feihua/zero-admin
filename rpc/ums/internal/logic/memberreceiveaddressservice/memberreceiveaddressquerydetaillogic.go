package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressQueryDetailLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 11:25
*/
type MemberReceiveAddressQueryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressQueryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressQueryDetailLogic {
	return &MemberReceiveAddressQueryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressQueryDetailLogic) MemberReceiveAddressQueryDetail(in *umsclient.MemberReceiveAddressQueryDetailReq) (*umsclient.MemberReceiveAddressQueryDetailResp, error) {
	q := query.UmsMemberReceiveAddress
	address, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.UserId), q.ID.Eq(in.AddressID)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员地址信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	resp := &umsclient.MemberReceiveAddressQueryDetailResp{
		Id:            address.ID,
		MemberId:      address.MemberID,
		MemberName:    address.MemberName,
		PhoneNumber:   address.PhoneNumber,
		DefaultStatus: address.DefaultStatus,
		PostCode:      address.PostCode,
		Province:      address.Province,
		City:          address.City,
		Region:        address.Region,
		DetailAddress: address.DetailAddress,
		CreateTime:    address.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:    address.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	logc.Infof(l.ctx, "查询会员地址信息,参数：%+v,响应：%+v", in, resp)

	return resp, nil
}
