package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCollectionListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:31
*/
type ProductCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCollectionListLogic {
	return &ProductCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCollectionList 查询会员收藏的商品
func (l *ProductCollectionListLogic) ProductCollectionList() (resp *types.ProductCollectionListResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	collectionList, _ := l.svcCtx.MemberProductCollectionService.MemberProductCollectionList(l.ctx, &umsclient.MemberProductCollectionListReq{
		Current:   1,
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
