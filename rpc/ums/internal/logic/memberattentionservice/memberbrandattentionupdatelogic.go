package memberattentionservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberBrandAttentionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberBrandAttentionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionUpdateLogic {
	return &MemberBrandAttentionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberBrandAttentionUpdateLogic) MemberBrandAttentionUpdate(in *umsclient.MemberBrandAttentionUpdateReq) (*umsclient.MemberBrandAttentionUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberBrandAttentionUpdateResp{}, nil
}
