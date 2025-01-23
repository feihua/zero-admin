package memberbrandattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberBrandAttentionLogic 添加会员关注品牌
/*
Author: LiuFeiHua
Date: 2024/6/11 14:18
*/
type AddMemberBrandAttentionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberBrandAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberBrandAttentionLogic {
	return &AddMemberBrandAttentionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberBrandAttention 添加会员关注品牌
func (l *AddMemberBrandAttentionLogic) AddMemberBrandAttention(in *umsclient.AddMemberBrandAttentionReq) (*umsclient.AddMemberBrandAttentionResp, error) {
	// 1.查询会员信息
	member, _ := query.UmsMember.WithContext(l.ctx).Where(query.UmsMember.ID.Eq(in.MemberId)).First()
	// 2.添加品牌关注
	err := query.UmsMemberBrandAttention.WithContext(l.ctx).Create(&model.UmsMemberBrandAttention{
		MemberID:       member.ID,
		MemberNickName: member.Nickname,
		MemberIcon:     member.Icon,
		BrandID:        in.BrandId,
		BrandName:      in.BrandName,
		BrandLogo:      in.BrandLogo,
		BrandCity:      in.BrandCity,
		CreateTime:     time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberBrandAttentionResp{}, nil
}
