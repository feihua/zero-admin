package memberaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberAddressListLogic 查询会员收货地址列表
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type QueryMemberAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberAddressListLogic {
	return &QueryMemberAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberAddressList 查询会员收货地址列表
func (l *QueryMemberAddressListLogic) QueryMemberAddressList(in *umsclient.QueryMemberAddressListReq) (*umsclient.QueryMemberAddressListResp, error) {
	memberAddress := query.UmsMemberAddress
	q := memberAddress.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(memberAddress.MemberID.Eq(in.MemberId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员收货地址列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员收货地址列表失败")
	}

	var list []*umsclient.MemberAddressListData

	for _, item := range result {
		list = append(list, &umsclient.MemberAddressListData{
			Id:            item.ID,                                 // 主键ID
			MemberId:      item.MemberID,                           // 会员ID
			ReceiverName:  item.ReceiverName,                       // 收货人姓名
			ReceiverPhone: item.ReceiverPhone,                      // 收货人电话
			Province:      item.Province,                           // 省份
			City:          item.City,                               // 城市
			District:      item.District,                           // 区县
			DetailAddress: item.DetailAddress,                      // 详细地址
			PostalCode:    item.PostalCode,                         // 邮政编码
			Tag:           item.Tag,                                // 地址标签：家、公司等
			IsDefault:     item.IsDefault,                          // 是否默认地址
			CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间
			IsDeleted:     item.IsDeleted,                          // 是否删除

		})
	}

	return &umsclient.QueryMemberAddressListResp{
		Total: count,
		List:  list,
	}, nil
}
