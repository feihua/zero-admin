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
	"gorm.io/gorm"
)

// QueryMemberAddressDetailLogic 查询会员收货地址详情
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type QueryMemberAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberAddressDetailLogic {
	return &QueryMemberAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberAddressDetail 查询会员收货地址详情
func (l *QueryMemberAddressDetailLogic) QueryMemberAddressDetail(in *umsclient.QueryMemberAddressDetailReq) (*umsclient.QueryMemberAddressDetailResp, error) {
	address := query.UmsMemberAddress
	item, err := address.WithContext(l.ctx).Where(address.ID.Eq(in.Id), address.MemberID.Eq(in.MemberId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员收货地址不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员收货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员收货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员收货地址异常")
	}

	data := &umsclient.QueryMemberAddressDetailResp{
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
	}

	return data, nil
}
