package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelAddLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/7 10:04
*/
type MemberLevelAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelAddLogic {
	return &MemberLevelAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberLevelAdd 添加会员等级
func (l *MemberLevelAddLogic) MemberLevelAdd(in *umsclient.MemberLevelAddReq) (*umsclient.MemberLevelAddResp, error) {
	db := l.svcCtx.DB
	q := query.Use(db)

	err := q.Transaction(func(tx *query.Query) error {
		levelDo := query.UmsMemberLevel.WithContext(l.ctx)
		if err := levelDo.Create(&model.UmsMemberLevel{
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
		}); err != nil {
			return err
		}

		//if _, err := levelDo.Where(query.UmsMemberLevel.ID.Eq(1)).Update(query.UmsMemberLevel.LevelName, "普通会员"); err != nil {
		//	return err
		//}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLevelAddResp{}, nil
}
