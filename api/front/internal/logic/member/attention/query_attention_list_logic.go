package attention

import (
	"context"

	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryAttentionListLogic 查询会员关注的列表
/*
Author: LiuFeiHua
Date: 2025/6/19 10:57
*/
type QueryAttentionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAttentionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAttentionListLogic {
	return &QueryAttentionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryAttentionList 查询会员关注的列表
func (l *QueryAttentionListLogic) QueryAttentionList() (resp *types.ListAttentionResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	attentionList, err := l.svcCtx.MemberBrandAttentionService.QueryMemberBrandAttentionList(l.ctx, &umsclient.QueryMemberBrandAttentionListReq{
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员关注的品牌失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	list := make([]types.ListAttentionData, 0)
	for _, detail := range attentionList.List {
		list = append(list, types.ListAttentionData{
			Id:             detail.Id,             //
			MemberId:       detail.MemberId,       // 会员id
			MemberNickName: detail.MemberNickName, // 会员姓名
			MemberIcon:     detail.MemberIcon,     // 会员头像
			BrandId:        detail.BrandId,        // 品牌id
			BrandName:      detail.BrandName,      // 品牌名称
			BrandLogo:      detail.BrandLogo,      // 品牌Logo
			BrandCity:      detail.BrandCity,      // 品牌所在城市
			CreateTime:     detail.CreateTime,     // 关注时间
		})
	}

	return &types.ListAttentionResp{
		Data:    list,
		Success: true,
		Code:    0,
		Message: "操作成功",
	}, nil
}
