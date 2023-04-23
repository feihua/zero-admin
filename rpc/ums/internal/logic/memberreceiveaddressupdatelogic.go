package logic

import (
	"context"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberReceiveAddressUpdateLogic) MemberReceiveAddressUpdate(in *umsclient.MemberReceiveAddressUpdateReq) (*umsclient.MemberReceiveAddressUpdateResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsMemberReceiveAddressModel.Update(l.ctx, &umsmodel.UmsMemberReceiveAddress{
		Id:            in.Id,
		MemberId:      in.MemberId,
		Name:          in.Name,
		PhoneNumber:   in.PhoneNumber,
		DefaultStatus: in.DefaultStatus,
		PostCode:      in.PostCode,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
		Address:       in.Address,
	})
	if err != nil {
		return nil, err
	}
	return &umsclient.MemberReceiveAddressUpdateResp{}, nil
}
