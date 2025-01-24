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
		LevelName:          in.LevelName,          // 等级名称
		GrowthPoint:        in.GrowthPoint,        // 成长点
		DefaultStatus:      in.DefaultStatus,      // 是否为默认等级：0->不是；1->是
		FreeFreightPoint:   in.FreeFreightPoint,   // 免运费标准
		CommentGrowthPoint: in.CommentGrowthPoint, // 每次评价获取的成长值
		IsFreeFreight:      in.IsFreeFreight,      // 是否有免邮特权
		IsSignIn:           in.IsSignIn,           // 是否有签到特权
		IsComment:          in.IsComment,          // 是否有评论获奖励特权
		IsPromotion:        in.IsPromotion,        // 是否有专享活动特权
		IsMemberPrice:      in.IsMemberPrice,      // 是否有会员价格特权
		IsBirthday:         in.IsBirthday,         // 是否有生日特权
		Remark:             in.Remark,             // 备注
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberLevelResp{}, nil
}
