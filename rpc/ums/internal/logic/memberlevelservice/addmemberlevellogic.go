package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberLevelLogic 添加会员等级表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:16
*/
type AddMemberLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLevelLogic {
	return &AddMemberLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberLevel 添加会员等级表
func (l *AddMemberLevelLogic) AddMemberLevel(in *umsclient.AddMemberLevelReq) (*umsclient.AddMemberLevelResp, error) {
	err := query.UmsMemberLevel.WithContext(l.ctx).Create(&model.UmsMemberLevel{
		LevelName:          in.LevelName,
		GrowthPoint:        in.GrowthPoint,
		DefaultStatus:      in.DefaultStatus,
		FreeFreightPoint:   in.FreeFreightPoint,
		CommentGrowthPoint: in.CommentGrowthPoint,
		IsFreeFreight:      in.IsFreeFreight,
		IsSignIn:           in.IsSignIn,
		IsComment:          in.IsComment,
		IsPromotion:        in.IsPromotion,
		IsMemberPrice:      in.IsMemberPrice,
		IsBirthday:         in.IsBirthday,
		Remark:             &in.Remark,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberLevelResp{}, nil
}
