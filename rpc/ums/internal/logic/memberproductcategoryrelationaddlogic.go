package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationAddLogic {
	return &MemberProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationAddLogic) MemberProductCategoryRelationAdd(in *ums.MemberProductCategoryRelationAddReq) (*ums.MemberProductCategoryRelationAddResp, error) {
	_, err := l.svcCtx.UmsMemberProductCategoryRelationModel.Insert(umsmodel.UmsMemberProductCategoryRelation{
		MemberId:          in.MemberId,
		ProductCategoryId: in.ProductCategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberProductCategoryRelationAddResp{}, nil
}
