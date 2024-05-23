package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 11:25
*/
type MemberReceiveAddressUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressUpdateLogic {
	return &MemberReceiveAddressUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReceiveAddressUpdate 修改会员收货地址
func (l *MemberReceiveAddressUpdateLogic) MemberReceiveAddressUpdate(in *umsclient.MemberReceiveAddressUpdateReq) (*umsclient.MemberReceiveAddressUpdateResp, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.UmsMemberReceiveAddress

		//如果更新的地址为默认地址,则需要把之前的默认地址去除默认标识
		addressDo := q.WithContext(l.ctx)
		if in.DefaultStatus == 1 {
			if _, err := addressDo.Where(q.MemberID.Eq(in.MemberId), q.DefaultStatus.Eq(1)).Update(q.DefaultStatus, 0); err != nil {
				return err
			}
		}

		if _, err := addressDo.Updates(&model.UmsMemberReceiveAddress{
			ID:            in.Id,
			MemberID:      in.MemberId,
			MemberName:    in.MemberName,
			PhoneNumber:   in.PhoneNumber,
			DefaultStatus: in.DefaultStatus,
			PostCode:      in.PostCode,
			Province:      in.Province,
			City:          in.City,
			Region:        in.Region,
			DetailAddress: in.DetailAddress,
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReceiveAddressUpdateResp{}, nil
}
