package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationDeleteLogic {
	return &MemberProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationDeleteLogic) MemberProductCategoryRelationDelete(in *ums.MemberProductCategoryRelationDeleteReq) (*ums.MemberProductCategoryRelationDeleteResp, error) {
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberProductCategoryRelationDeleteResp{}, nil
}
