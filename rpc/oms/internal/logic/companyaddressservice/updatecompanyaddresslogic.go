package companyaddressservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressLogic 更新公司收发货地址表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:11
*/
type UpdateCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressLogic {
	return &UpdateCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCompanyAddress 更新公司收发货地址表
func (l *UpdateCompanyAddressLogic) UpdateCompanyAddress(in *omsclient.UpdateCompanyAddressReq) (*omsclient.UpdateCompanyAddressResp, error) {

	err := query.Q.Transaction(func(tx *query.Query) error {
		q := tx.OmsCompanyAddress

		_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return errors.New("公司收发货地址不存在")
		case err != nil:
			logc.Errorf(l.ctx, "查询公司收发货地址异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
			return errors.New("查询公司收发货地址异常")
		}

		count, err := q.WithContext(l.ctx).Where(q.AddressName.Eq(in.AddressName), q.ID.Neq(in.Id)).Count()
		if count > 0 {
			return errors.New(fmt.Sprintf("更新公司收发货地址表失败,地址名称已存在"))
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

		err = l.svcCtx.DB.Model(&model.OmsCompanyAddress{}).WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Save(&model.OmsCompanyAddress{
			ID:            in.Id,
			AddressName:   in.AddressName,   // 地址名称
			SendStatus:    in.SendStatus,    // 默认发货地址：0->否；1->是
			ReceiveStatus: in.ReceiveStatus, // 是否默认收货地址：0->否；1->是
			Name:          in.Name,          // 收发货人姓名
			Phone:         in.Phone,         // 收货人电话
			Province:      in.Province,      // 省/直辖市
			City:          in.City,          // 市
			Region:        in.Region,        // 区
			DetailAddress: in.DetailAddress, // 详细地址
			UpdateBy:      in.UpdateBy,      // 创建者
		}).Error

		return err
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressResp{}, nil
}
