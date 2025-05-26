package companyaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCompanyAddressDetailLogic 查询公司收发货地址详情
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:43
*/
type QueryCompanyAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCompanyAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressDetailLogic {
	return &QueryCompanyAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCompanyAddressDetail 查询公司收发货地址详情
func (l *QueryCompanyAddressDetailLogic) QueryCompanyAddressDetail(in *omsclient.QueryCompanyAddressDetailReq) (*omsclient.QueryCompanyAddressDetailResp, error) {
	item, err := query.OmsCompanyAddress.WithContext(l.ctx).Where(query.OmsCompanyAddress.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "公司收发货地址不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("公司收发货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询公司收发货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询公司收发货地址异常")
	}

	data := &omsclient.QueryCompanyAddressDetailResp{
		Id:            item.ID,                                          // 主键ID
		AddressName:   item.AddressName,                                 // 地址名称
		Name:          item.Name,                                        // 收发货人姓名
		Phone:         item.Phone,                                       // 收货人电话
		Province:      item.Province,                                    // 省/直辖市
		City:          item.City,                                        // 市
		Region:        item.Region,                                      // 区
		DetailAddress: item.DetailAddress,                               // 详细地址
		SendStatus:    item.SendStatus,                                  // 默认发货地址：0->否；1->是
		ReceiveStatus: item.ReceiveStatus,                               // 默认收货地址：0->否；1->是
		CreateBy:      item.CreateBy,                                    // 创建者
		CreateTime:    time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:      pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:    time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
