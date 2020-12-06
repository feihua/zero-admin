package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationUpdateLogic {
	return &MemberProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationUpdateLogic) MemberProductCategoryRelationUpdate(in *ums.MemberProductCategoryRelationUpdateReq) (*ums.MemberProductCategoryRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberProductCategoryRelationUpdateResp{}, nil
}
