package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelUpdateLogic {
	return MemberLevelUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLevelUpdateLogic) MemberLevelUpdate(req types.UpdateMemberLevelReq) (*types.UpdateMemberLevelResp, error) {
	_, err := l.svcCtx.Ums.MemberLevelUpdate(l.ctx, &umsclient.MemberLevelUpdateReq{
		Id:                    req.Id,
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

	return &types.UpdateMemberLevelResp{
		Code:    "000000",
		Message: "",
	}, nil
}
