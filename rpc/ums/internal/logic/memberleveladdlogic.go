package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelAddLogic {
	return &MemberLevelAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(in *ums.MemberLevelAddReq) (*ums.MemberLevelAddResp, error) {
	_, err := l.svcCtx.UmsMemberLevelModel.Insert(umsmodel.UmsMemberLevel{
		Name:                  in.Name,
		GrowthPoint:           in.GrowthPoint,
		DefaultStatus:         in.DefaultStatus,
		FreeFreightPoint:      float64(in.FreeFreightPoint),
		CommentGrowthPoint:    in.CommentGrowthPoint,
		PriviledgeFreeFreight: in.PriviledgeFreeFreight,
		PriviledgeSignIn:      in.PriviledgeSignIn,
		PriviledgeComment:     in.PriviledgeComment,
		PriviledgePromotion:   in.PriviledgePromotion,
		PriviledgeMemberPrice: in.PriviledgeMemberPrice,
		PriviledgeBirthday:    in.PriviledgeBirthday,
		Note:                  in.Note,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberLevelAddResp{}, nil
}
