package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressDeleteLogic {
	return &MemberAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressDeleteLogic) MemberAddressDelete(req *types.DeleteMemberAddressReq) (resp *types.DeleteMemberAddressResp, err error) {
	_, _ = l.svcCtx.Ums.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Ids:      req.Ids,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	return &types.DeleteMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
