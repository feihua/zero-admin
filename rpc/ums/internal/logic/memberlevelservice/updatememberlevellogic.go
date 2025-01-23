package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLevelLogic 更新会员等级表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:17
*/
type UpdateMemberLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelLogic {
	return &UpdateMemberLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberLevel 更新会员等级表
func (l *UpdateMemberLevelLogic) UpdateMemberLevel(in *umsclient.UpdateMemberLevelReq) (*umsclient.UpdateMemberLevelResp, error) {
	_, err := query.UmsMemberLevel.WithContext(l.ctx).Updates(&model.UmsMemberLevel{
		ID:                 in.Id,
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
		Remark:             in.Remark,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberLevelResp{}, nil
}
