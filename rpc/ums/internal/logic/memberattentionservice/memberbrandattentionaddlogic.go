package memberattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberBrandAttentionAddLogic 品牌关注相关
/*
Author: LiuFeiHua
Date: 2023/12/4 16:51
*/
type MemberBrandAttentionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberBrandAttentionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionAddLogic {
	return &MemberBrandAttentionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberBrandAttentionAdd 添加品牌关注
func (l *MemberBrandAttentionAddLogic) MemberBrandAttentionAdd(in *umsclient.MemberBrandAttentionAddReq) (*umsclient.MemberBrandAttentionAddResp, error) {
	//1.查询会员信息
	member, _ := query.UmsMember.WithContext(l.ctx).Where(query.UmsMember.ID.Eq(in.MemberId)).First()
	//2.添加品牌关注
	err := query.UmsMemberBrandAttention.WithContext(l.ctx).Create(&model.UmsMemberBrandAttention{
		MemberID:       member.ID,
		MemberNickName: member.Nickname,
		MemberIcon:     *member.Icon,
		BrandID:        in.BrandId,
		BrandName:      in.BrandName,
		BrandLogo:      in.BrandLogo,
		BrandCity:      &in.BrandCity,
		CreateTime:     time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberBrandAttentionAddResp{}, nil
}
