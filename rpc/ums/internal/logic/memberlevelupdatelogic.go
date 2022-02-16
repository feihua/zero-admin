package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelUpdateLogic {
	return &MemberLevelUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelUpdateLogic) MemberLevelUpdate(in *ums.MemberLevelUpdateReq) (*ums.MemberLevelUpdateResp, error) {
	err := l.svcCtx.UmsMemberLevelModel.Update(umsmodel.UmsMemberLevel{
		Id:                    in.Id,
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

	return &ums.MemberLevelUpdateResp{}, nil
}
