package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductCollectionLogic 添加商品收藏
/*
Author: LiuFeiHua
Date: 2024/5/16 13:45
*/
type AddProductCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCollectionLogic {
	return &AddProductCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductCollection 添加商品收藏
func (l *AddProductCollectionLogic) AddProductCollection(req *types.AddProductCollectionReq) (resp *types.AddProductCollectionResp, err error) {
	//从token中获取会员id
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	member, _ := l.svcCtx.MemberService.QueryMemberDetail(l.ctx, &umsclient.QueryMemberDetailReq{Id: memberId})

	//查询是否已经收藏
	collectionList, _ := l.svcCtx.MemberProductCollectionService.QueryMemberProductCollectionList(l.ctx, &umsclient.QueryMemberProductCollectionListReq{
		PageNum:   1,
		PageSize:  10,
		MemberId:  memberId,
		ProductId: req.ProductId,
	})

	//如果查询不到收藏记录,则添加
	if len(collectionList.List) == 0 {
		_, err = l.svcCtx.MemberProductCollectionService.AddMemberProductCollection(l.ctx, &umsclient.AddMemberProductCollectionReq{
			MemberId:        member.Id,
			MemberNickName:  member.Nickname,
			MemberIcon:      member.Icon,
			ProductId:       req.ProductId,
			ProductName:     req.ProductName,
			ProductPic:      req.ProductPic,
			ProductSubTitle: req.ProductSubTitle,
			ProductPrice:    req.ProductPrice,
		})

		if err != nil {
			return nil, err
		}
	}

	return &types.AddProductCollectionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
