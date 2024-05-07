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
		ID:                    in.Id,
		Name:                  in.Name,
		GrowthPoint:           in.GrowthPoint,
		DefaultStatus:         in.DefaultStatus,
		FreeFreightPoint:      float64(in.FreeFreightPoint),
		CommentGrowthPoint:    in.CommentGrowthPoint,
		PriviledgeFreeFreight: in.PriviledgeFreeFreight,
		PriviledgeSignIn:      in.PriviledgeSignIn,
		PriviledgeComment:     in.PriviledgeComment,
		PriviledgePromotion:   in.PriviledgePromotion,
		PriviledgeMemberPrice: in.PriviledgeMemberPrice,
		PriviledgeBirthday:    in.PriviledgeBirthday,
		Note:                  &in.Note,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLevelUpdateResp{}, nil
}
