package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *MemberTagAddLogic) MemberTagAdd(in *ums.MemberTagAddReq) (*ums.MemberTagAddResp, error) {
	_, err := l.svcCtx.UmsMemberTagModel.Insert(umsmodel.UmsMemberTag{
		Name:              in.Name,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: float64(in.FinishOrderAmount),
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTagAddResp{}, nil
}
