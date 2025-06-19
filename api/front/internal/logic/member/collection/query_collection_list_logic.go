package collection

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

// QueryCollectionListLogic 查询会员收藏的商品
/*
Author: LiuFeiHua
Date: 2025/6/19 10:58
*/
type QueryCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCollectionListLogic {
	return &QueryCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCollectionList 查询会员收藏的商品
func (l *QueryCollectionListLogic) QueryCollectionList() (resp *types.CollectionListResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	collectionList, err := l.svcCtx.MemberProductCollectionService.QueryMemberProductCollectionList(l.ctx, &umsclient.QueryMemberProductCollectionListReq{
		PageNum:   1,
		PageSize:  100,
		MemberId:  memberId,
		ProductId: 0,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员收藏的商品失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []types.CollectionList

	for _, detail := range collectionList.List {
		list = append(list, types.CollectionList{
			Id:              detail.Id,              //
			MemberId:        detail.MemberId,        // 会员id
			MemberNickName:  detail.MemberNickName,  // 会员姓名
			MemberIcon:      detail.MemberIcon,      // 会员头像
			ProductId:       detail.ProductId,       // 商品id
			ProductName:     detail.ProductName,     // 商品名称
			ProductPic:      detail.ProductPic,      // 商品图片
			ProductSubTitle: detail.ProductSubTitle, // 商品标题
			ProductPrice:    detail.ProductPrice,    // 商品价格
			CreateTime:      detail.CreateTime,      // 收藏时间
		})
	}

	return &types.CollectionListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
