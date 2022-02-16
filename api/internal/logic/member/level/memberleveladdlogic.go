package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员等级信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员等级失败")
	}

	return &types.AddMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
