package memberreceiveaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberReceiveAddressDetailLogic 查询会员收货地址详情
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

// QueryMemberReceiveAddressDetail 查询会员收货地址详情
func (l *QueryMemberReceiveAddressDetailLogic) QueryMemberReceiveAddressDetail(in *umsclient.QueryMemberReceiveAddressDetailReq) (*umsclient.QueryMemberReceiveAddressDetailResp, error) {
	q := query.UmsMemberReceiveAddress
	item, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId), q.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员收货地址详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员收货地址详情失败")
	}

	resp := &umsclient.QueryMemberReceiveAddressDetailResp{
		Id:            item.ID,                                 //
		MemberId:      item.MemberID,                           // 会员id
		MemberName:    item.MemberName,                         // 收货人名称
		PhoneNumber:   item.PhoneNumber,                        // 收货人电话
		DefaultStatus: item.DefaultStatus,                      // 是否为默认
		PostCode:      item.PostCode,                           // 邮政编码
		Province:      item.Province,                           // 省份/直辖市
		City:          item.City,                               // 城市
		Region:        item.Region,                             // 区
		DetailAddress: item.DetailAddress,                      // 详细地址(街道)
		CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return resp, nil
}
