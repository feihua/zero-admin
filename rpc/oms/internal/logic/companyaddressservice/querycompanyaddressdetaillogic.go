package companyaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCompanyAddressDetailLogic 查询公司收发货地址表详情
/*
Author: 刘飞华
Date: 2025/02/06 13:53:02
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

// QueryCompanyAddressDetail 查询公司收发货地址表详情
func (l *QueryCompanyAddressDetailLogic) QueryCompanyAddressDetail(in *omsclient.QueryCompanyAddressDetailReq) (*omsclient.QueryCompanyAddressDetailResp, error) {
	item, err := query.OmsCompanyAddress.WithContext(l.ctx).Where(query.OmsCompanyAddress.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("公司收发货地址不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询公司收发货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询公司收发货地址异常")
	}

	data := &omsclient.QueryCompanyAddressDetailResp{
		Id:            item.ID,                                 //
		AddressName:   item.AddressName,                        // 地址名称
		SendStatus:    item.SendStatus,                         // 默认发货地址：0->否；1->是
		ReceiveStatus: item.ReceiveStatus,                      // 是否默认收货地址：0->否；1->是
		Name:          item.Name,                               // 收发货人姓名
		Phone:         item.Phone,                              // 收货人电话
		Province:      item.Province,                           // 省/直辖市
		City:          item.City,                               // 市
		Region:        item.Region,                             // 区
		DetailAddress: item.DetailAddress,                      // 详细地址
		CreateBy:      item.CreateBy,                           // 创建者
		CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:      item.UpdateBy,                           // 更新者
		UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间

	}

	logc.Infof(l.ctx, "查询公司收发货地址表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
