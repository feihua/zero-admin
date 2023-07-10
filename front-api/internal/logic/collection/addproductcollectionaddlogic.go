package collection

import (
	"context"
	"time"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCollectionAddLogic {
	return &AddProductCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductCollectionAdd 会员收藏商品
func (l *AddProductCollectionAddLogic) AddProductCollectionAdd(req *types.AddProductCollectionReq) (resp *types.AddProductCollectionResp, err error) {
	//从token中获取会员id
	memberId := l.ctx.Value("memberId").(int64)
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})

	//查询是否已经收藏
	collectionList, _ := l.svcCtx.MemberProductCollectionService.MemberProductCollectionList(l.ctx, &umsclient.MemberProductCollectionListReq{
		Current:   1,
		PageSize:  10,
		MemberId:  memberId,
		ProductId: req.ProductId,
	})

	//如果查询不到收藏记录,则添加
	if len(collectionList.List) == 0 {
		_, err = l.svcCtx.MemberProductCollectionService.MemberProductCollectionAdd(l.ctx, &umsclient.MemberProductCollectionAddReq{
			MemberId:        member.Id,
			MemberNickName:  member.Nickname,
			MemberIcon:      member.Icon,
			ProductId:       req.ProductId,
			ProductName:     req.ProductName,
			ProductPic:      req.ProductPic,
			ProductSubTitle: req.ProductSubTitle,
			ProductPrice:    req.ProductPrice,
			CreateTime:      time.Now().Format("2006-01-02 15:04:05"),
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
