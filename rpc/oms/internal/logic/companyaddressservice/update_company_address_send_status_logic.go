package companyaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCompanyAddressSendStatusLogic 更新公司收发货地址状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:35
*/
type UpdateCompanyAddressSendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCompanyAddressSendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyAddressSendStatusLogic {
	return &UpdateCompanyAddressSendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCompanyAddressSendStatus 更新公司收发货地址状态
func (l *UpdateCompanyAddressSendStatusLogic) UpdateCompanyAddressSendStatus(in *omsclient.UpdateCompanyAddressStatusReq) (*omsclient.UpdateCompanyAddressStatusResp, error) {
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
		if in.Status == 1 {
			_, err = q.WithContext(l.ctx).Where(q.SendStatus.Eq(1)).UpdateSimple(q.SendStatus.Value(0), q.UpdateBy.Value(in.UpdateBy))
			if err != nil {
				return err
			}
		}

		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.SendStatus.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))
		return err
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCompanyAddressStatusResp{}, nil
}
