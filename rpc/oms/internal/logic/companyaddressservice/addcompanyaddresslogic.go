package companyaddressservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCompanyAddressLogic 添加公司收发货地址
/*
Author: LiuFeiHua
Date: 2024/6/12 10:10
*/
type AddCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompanyAddressLogic {
	return &AddCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCompanyAddress 添加公司收发货地址
func (l *AddCompanyAddressLogic) AddCompanyAddress(in *omsclient.AddCompanyAddressReq) (*omsclient.AddCompanyAddressResp, error) {

	err := query.Q.Transaction(func(tx *query.Query) error {
		q := tx.OmsCompanyAddress
		count, err := q.WithContext(l.ctx).Where(q.AddressName.Eq(in.AddressName)).Count()
		if count > 0 {
			return errors.New(fmt.Sprintf("添加公司收发货地址失败,地址名称已存在"))
		}

		if in.ReceiveStatus == 1 {
			_, err = q.WithContext(l.ctx).Where(q.ReceiveStatus.Eq(1)).Update(q.ReceiveStatus, 0)
			if err != nil {
				return err
			}
		}
		if in.SendStatus == 1 {
			_, err = q.WithContext(l.ctx).Where(q.SendStatus.Eq(1)).Update(q.SendStatus, 0)
			if err != nil {
				return err
			}
		}

		return q.WithContext(l.ctx).Create(&model.OmsCompanyAddress{
			AddressName:   in.AddressName,   // 地址名称
			SendStatus:    in.SendStatus,    // 默认发货地址：0->否；1->是
			ReceiveStatus: in.ReceiveStatus, // 是否默认收货地址：0->否；1->是
			Name:          in.Name,          // 收发货人姓名
			Phone:         in.Phone,         // 收货人电话
			Province:      in.Province,      // 省/直辖市
			City:          in.City,          // 市
			Region:        in.Region,        // 区
			DetailAddress: in.DetailAddress, // 详细地址
			CreateBy:      in.CreateBy,      // 创建者
		})

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加公司收发货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加公司收发货地址失败")
	}

	return &omsclient.AddCompanyAddressResp{}, nil
}
