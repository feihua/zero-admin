package attention

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AttentionListLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 17:15
*/
type AttentionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttentionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttentionListLogic {
	return &AttentionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AttentionList 查询会员关注的列表
func (l *AttentionListLogic) AttentionList() (resp *types.ListAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	var attentionList, _ = l.svcCtx.MemberAttentionService.MemberBrandAttentionList(l.ctx, &umsclient.MemberBrandAttentionListReq{
		MemberId: memberId,
	})

	var list []types.ListAttentionData

	for _, item := range attentionList.List {
		list = append(list, types.ListAttentionData{
			Id:             item.Id,
			MemberId:       item.MemberId,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			BrandId:        item.BrandId,
			BrandName:      item.BrandName,
			BrandLogo:      item.BrandLogo,
			BrandCity:      item.BrandCity,
			CreateTime:     item.CreateTime,
		})
	}

	return &types.ListAttentionResp{
		Data:    list,
		Success: true,
		Code:    0,
		Message: "操作成功",
	}, nil
}
