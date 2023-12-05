package memberattentionservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	member, _ := l.svcCtx.UmsMemberModel.FindOne(l.ctx, in.MemberId)
	//2.添加品牌关注
	_, err := l.svcCtx.UmsMemberBrandAttentionModel.Insert(l.ctx, &umsmodel.UmsMemberBrandAttention{
		MemberId:       member.Id,
		MemberNickName: member.Nickname,
		MemberIcon:     member.Icon.String,
		BrandId:        in.BrandId,
		BrandName:      in.BrandName,
		BrandLogo:      in.BrandLogo,
		BrandCity: sql.NullString{
			String: in.BrandCity,
			Valid:  true,
		},
		CreateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberBrandAttentionAddResp{}, nil
}
