package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCollectionListLogic 查询会员收藏的商品
/*
Author: LiuFeiHua
Date: 2024/5/16 13:44
*/
type QueryProductCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCollectionListLogic {
	return &QueryProductCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductCollectionList 查询会员收藏的商品
func (l *QueryProductCollectionListLogic) QueryProductCollectionList() (resp *types.ProductCollectionListResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	collectionList, _ := l.svcCtx.MemberProductCollectionService.QueryMemberProductCollectionList(l.ctx, &umsclient.QueryMemberProductCollectionListReq{
		PageNum:   1,
		PageSize:  100,
		MemberId:  memberId,
		ProductId: 0,
	})

	var list []types.ProductCollectionList

	for _, member := range collectionList.List {
		list = append(list, types.ProductCollectionList{
			Id:              member.Id,
			MemberId:        member.MemberId,
			MemberNickName:  member.MemberNickName,
			MemberIcon:      member.MemberIcon,
			ProductId:       member.ProductId,
			ProductName:     member.ProductName,
			ProductPic:      member.ProductPic,
			ProductSubTitle: member.ProductSubTitle,
			ProductPrice:    member.ProductPrice,
			CreateTime:      member.CreateTime,
		})
	}

	return &types.ProductCollectionListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
