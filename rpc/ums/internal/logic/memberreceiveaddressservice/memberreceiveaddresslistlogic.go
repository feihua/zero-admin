package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:17
*/
type MemberReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressListLogic {
	return &MemberReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReceiveAddressList 查询会员收货地址列表
func (l *MemberReceiveAddressListLogic) MemberReceiveAddressList(in *umsclient.MemberReceiveAddressListReq) (*umsclient.MemberReceiveAddressListResp, error) {
	q := query.UmsMemberReceiveAddress.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberLoginLog.MemberID.Eq(in.MemberId))
	}
	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员地址列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberReceiveAddressListData
	for _, address := range result {

		list = append(list, &umsclient.MemberReceiveAddressListData{
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
		})
	}

	logc.Infof(l.ctx, "查询会员地址列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberReceiveAddressListResp{
		Total: count,
		List:  list,
	}, nil

}
