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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员等级信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员等级失败")
	}

	return &types.UpdateMemberLevelResp{
		Code:    "000000",
		Message: "更新会员等级成功",
	}, nil
}
