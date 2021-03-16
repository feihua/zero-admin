package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelAddLogic {
	return MemberLevelAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelAddLogic) MemberLevelAdd(req types.AddMemberLevelReq) (*types.AddMemberLevelResp, error) {
	_, err := l.svcCtx.Ums.MemberLevelAdd(l.ctx, &umsclient.MemberLevelAddReq{
		Name:                  req.Name,
		GrowthPoint:           req.GrowthPoint,
		DefaultStatus:         req.DefaultStatus,
		FreeFreightPoint:      int64(req.FreeFreightPoint),
		CommentGrowthPoint:    req.CommentGrowthPoint,
		PriviledgeFreeFreight: req.PriviledgeFreeFreight,
		PriviledgeSignIn:      req.PriviledgeSignIn,
		PriviledgeComment:     req.PriviledgeComment,
		PriviledgePromotion:   req.PriviledgePromotion,
		PriviledgeMemberPrice: req.PriviledgeMemberPrice,
		PriviledgeBirthday:    req.PriviledgeBirthday,
		Note:                  req.Note,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberLevelResp{
		Code:    "000000",
		Message: "",
	}, nil
}
