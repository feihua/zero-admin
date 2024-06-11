package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/logic/common"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberReceiveAddressDetailLogic 查询会员收货地址表详情
/*
Author: LiuFeiHua
Date: 2024/6/11 14:05
*/
type QueryMemberReceiveAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReceiveAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReceiveAddressDetailLogic {
	return &QueryMemberReceiveAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberReceiveAddressDetail 查询会员收货地址表详情
func (l *QueryMemberReceiveAddressDetailLogic) QueryMemberReceiveAddressDetail(in *umsclient.QueryMemberReceiveAddressDetailReq) (*umsclient.QueryMemberReceiveAddressDetailResp, error) {
	q := query.UmsMemberReceiveAddress
	address, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId), q.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员地址信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	resp := &umsclient.QueryMemberReceiveAddressDetailResp{
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
		UpdateTime:    common.TimeToString(address.UpdateTime),
	}

	return resp, nil
}
