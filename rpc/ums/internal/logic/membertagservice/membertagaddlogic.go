package membertagservicelogic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagAddLogic {
	return &MemberTagAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagAddLogic) MemberTagAdd(in *umsclient.MemberTagAddReq) (*umsclient.MemberTagAddResp, error) {
	_, err := l.svcCtx.UmsMemberTagModel.Insert(l.ctx, &umsmodel.UmsMemberTag{
		Name:              in.Name,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: float64(in.FinishOrderAmount),
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagAddResp{}, nil
}
