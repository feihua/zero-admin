package level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberLevelLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:37
*/
type AddMemberLevelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLevelLogic {
	return &AddMemberLevelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberLevel 添加会员等级
func (l *AddMemberLevelLogic) AddMemberLevel(req *types.AddMemberLevelReq) (resp *types.AddMemberLevelResp, err error) {
	_, err = l.svcCtx.MemberLevelService.MemberLevelAdd(l.ctx, &umsclient.MemberLevelAddReq{
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
		logc.Errorf(l.ctx, "添加会员等级信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员等级失败")
	}

	return &types.AddMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
