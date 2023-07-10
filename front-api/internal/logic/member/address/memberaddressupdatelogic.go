package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressUpdateLogic {
	return &MemberAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressUpdateLogic) MemberAddressUpdate(req *types.UpdateMemberAddressReq) (resp *types.UpdateMemberAddressResp, err error) {
	_, err = l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.Id,
		MemberId:      l.ctx.Value("memberId").(int64),
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
		Code:    0,
		Message: "操作成功",
	}, nil
}
