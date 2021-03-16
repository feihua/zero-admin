package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressUpdateLogic {
	return MemberAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressUpdateLogic) MemberAddressUpdate(req types.UpdateMemberAddressReq) (*types.UpdateMemberAddressResp, error) {
	_, err := l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.Id,
		MemberId:      req.MemberId,
		Name:          req.Name,
		PhoneNumber:   req.PhoneNumber,
		DefaultStatus: req.DefaultStatus,
		PostCode:      req.PostCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
