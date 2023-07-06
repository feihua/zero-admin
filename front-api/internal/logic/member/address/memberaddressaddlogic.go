package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressAddLogic {
	return &MemberAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(req *types.AddMemberAddressReq) (resp *types.AddMemberAddressResp, err error) {
	_, err = l.svcCtx.Ums.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
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

	return &types.AddMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
