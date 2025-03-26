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
Date: 2024/5/16 11:04
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

	var list []types.ListAttentionData

	for _, item := range attentionList.List {
		list = append(list, types.ListAttentionData{
			Id:         item.Id,
			BrandId:    item.BrandId,
			BrandName:  item.BrandName,
			BrandLogo:  item.BrandLogo,
			BrandCity:  item.BrandCity,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListAttentionResp{
		Data:    list,
		Success: true,
		Code:    0,
		Message: "操作成功",
	}, nil
}
