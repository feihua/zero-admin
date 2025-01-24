package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberReceiveAddressLogic 更新会员收货地址表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:05
*/
type UpdateMemberReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReceiveAddressLogic {
	return &UpdateMemberReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberReceiveAddress 更新会员收货地址表
func (l *UpdateMemberReceiveAddressLogic) UpdateMemberReceiveAddress(in *umsclient.UpdateMemberReceiveAddressReq) (*umsclient.UpdateMemberReceiveAddressResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.UmsMemberReceiveAddress

		// 如果更新的地址为默认地址,则需要把之前的默认地址去除默认标识
		addressDo := q.WithContext(l.ctx)
		if in.DefaultStatus == 1 {
			if _, err := addressDo.Where(q.MemberID.Eq(in.MemberId), q.DefaultStatus.Eq(1)).Update(q.DefaultStatus, 0); err != nil {
				return err
			}
		}

		if _, err := addressDo.Updates(&model.UmsMemberReceiveAddress{
			ID:            in.Id,
			MemberID:      in.MemberId,      // 会员id
			MemberName:    in.MemberName,    // 收货人名称
			PhoneNumber:   in.PhoneNumber,   // 收货人电话
			DefaultStatus: in.DefaultStatus, // 是否为默认
			PostCode:      in.PostCode,      // 邮政编码
			Province:      in.Province,      // 省份/直辖市
			City:          in.City,          // 城市
			Region:        in.Region,        // 区
			DetailAddress: in.DetailAddress, // 详细地址(街道)
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberReceiveAddressResp{}, nil
}
