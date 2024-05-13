package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelUpdateLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:38
*/
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

// MemberLevelUpdate 更新会员等级
func (l *MemberLevelUpdateLogic) MemberLevelUpdate(req types.UpdateMemberLevelReq) (*types.UpdateMemberLevelResp, error) {
	_, err := l.svcCtx.MemberLevelService.MemberLevelUpdate(l.ctx, &umsclient.MemberLevelUpdateReq{
		Id:                 req.Id,
		LevelName:          req.Name,
		GrowthPoint:        req.GrowthPoint,
		DefaultStatus:      req.DefaultStatus,
		FreeFreightPoint:   req.FreeFreightPoint,
		CommentGrowthPoint: req.CommentGrowthPoint,
		IsFreeFreight:      req.IsFreeFreight,
		IsSignIn:           req.IsSignIn,
		IsComment:          req.IsComment,
		IsPromotion:        req.IsPromotion,
		IsMemberPrice:      req.IsMemberPrice,
		IsBirthday:         req.IsBirthday,
		Remark:             req.Note,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员等级信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员等级失败")
	}

	return &types.UpdateMemberLevelResp{
		Code:    "000000",
		Message: "更新会员等级成功",
	}, nil
}
