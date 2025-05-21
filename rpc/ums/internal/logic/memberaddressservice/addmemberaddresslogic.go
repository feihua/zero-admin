package memberaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberAddressLogic 添加会员收货地址
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type AddMemberAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberAddressLogic {
	return &AddMemberAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberAddress 添加会员收货地址
func (l *AddMemberAddressLogic) AddMemberAddress(in *umsclient.AddMemberAddressReq) (*umsclient.AddMemberAddressResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.UmsMemberAddress

		// 如果新增的地址为默认地址,则需要把之前的默认地址去除默认标识
		addressDo := q.WithContext(l.ctx)
		if in.IsDefault == 1 {
			if _, err := addressDo.Where(q.MemberID.Eq(in.MemberId), q.IsDefault.Eq(1)).Update(q.IsDefault, 0); err != nil {
				return err
			}
		}

		if err := addressDo.Create(&model.UmsMemberAddress{
			MemberID:      in.MemberId,      // 会员ID
			ReceiverName:  in.ReceiverName,  // 收货人姓名
			ReceiverPhone: in.ReceiverPhone, // 收货人电话
			Province:      in.Province,      // 省份
			City:          in.City,          // 城市
			District:      in.District,      // 区县
			DetailAddress: in.DetailAddress, // 详细地址
			PostalCode:    in.PostalCode,    // 邮政编码
			Tag:           in.Tag,           // 地址标签：家、公司等
			IsDefault:     in.IsDefault,     // 是否默认地址
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员收货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加会员收货地址失败")
	}

	return &umsclient.AddMemberAddressResp{}, nil
}
