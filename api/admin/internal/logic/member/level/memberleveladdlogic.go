package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.MemberLevelService.MemberLevelAdd(l.ctx, &umsclient.MemberLevelAddReq{
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员等级信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员等级失败")
	}

	return &types.AddMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
