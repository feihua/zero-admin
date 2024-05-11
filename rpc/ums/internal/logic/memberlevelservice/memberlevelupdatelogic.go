package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelUpdateLogic 会员等级
type MemberLevelUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelUpdateLogic {
	return &MemberLevelUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberLevelUpdate 更新会员等级
func (l *MemberLevelUpdateLogic) MemberLevelUpdate(in *umsclient.MemberLevelUpdateReq) (*umsclient.MemberLevelUpdateResp, error) {
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
		Remark:             &in.Remark,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLevelUpdateResp{}, nil
}
