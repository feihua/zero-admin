package collection

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductCollectionListLogic) ProductCollectionList(req *types.ProductCollectionListReq) (resp *types.ProductCollectionListResp, err error) {
	collectionList, _ := l.svcCtx.Ums.MemberProductCollectionList(l.ctx, &umsclient.MemberProductCollectionListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		MemberId:  l.ctx.Value("memberId").(int64),
		ProductId: req.ProductId,
	})

	var list []*types.ProductCollectionList

	for _, member := range collectionList.List {
		list = append(list, &types.ProductCollectionList{
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
		Data: types.ProductCollectionListData{
			Total: collectionList.Total,
			Pages: collectionList.Total,
			Limit: req.PageSize,
			Page:  req.Current,
			List:  nil,
		},
	}, nil
}
