package memberbrandattentionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
		MemberID:       in.MemberId,       // 会员id
		MemberNickName: member.MemberName, // 会员姓名
		MemberIcon:     member.Icon,       // 会员头像
		BrandID:        in.BrandId,        // 品牌id
		BrandName:      in.BrandName,      // 品牌名称
		BrandLogo:      in.BrandLogo,      // 品牌Logo
		BrandCity:      in.BrandCity,      // 品牌所在城市
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员关注品牌失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加会员关注品牌失败")
	}

	return &umsclient.AddMemberBrandAttentionResp{}, nil
}
