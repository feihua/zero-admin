package logic

import (
	"context"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReceiveAddressAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressAddLogic {
	return &MemberReceiveAddressAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressAddLogic) MemberReceiveAddressAdd(in *umsclient.MemberReceiveAddressAddReq) (*umsclient.MemberReceiveAddressAddResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UmsMemberReceiveAddressModel.Insert(l.ctx, &umsmodel.UmsMemberReceiveAddress{
		MemberId:      in.MemberId,
		Name:          in.Name,
		PhoneNumber:   in.PhoneNumber,
		DefaultStatus: in.DefaultStatus,
		PostCode:      in.PostCode,
		Province:      in.Province,
		City:          in.City,
		Region:        in.Region,
		DetailAddress: in.DetailAddress,
	})
	if err != nil {
		return nil, err
	}
	return &umsclient.MemberReceiveAddressAddResp{}, nil
}
