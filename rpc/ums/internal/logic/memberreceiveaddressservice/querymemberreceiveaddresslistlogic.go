package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberReceiveAddressListLogic 查询会员收货地址表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:04
*/
type QueryMemberReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReceiveAddressListLogic {
	return &QueryMemberReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberReceiveAddressList 查询会员收货地址表列表
func (l *QueryMemberReceiveAddressListLogic) QueryMemberReceiveAddressList(in *umsclient.QueryMemberReceiveAddressListReq) (*umsclient.QueryMemberReceiveAddressListResp, error) {
	q := query.UmsMemberReceiveAddress.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberLoginLog.MemberID.Eq(in.MemberId))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

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
			UpdateTime:    time_util.TimeToString(address.UpdateTime),
		})
	}

	return &umsclient.QueryMemberReceiveAddressListResp{
		Total: count,
		List:  list,
	}, nil

}
