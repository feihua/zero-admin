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

// AddCollectionLogic 添加商品收藏
/*
Author: LiuFeiHua
Date: 2025/6/19 10:58
*/
type AddCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddCollection 添加商品收藏
func (l *AddCollectionLogic) AddCollection(req *types.CollectionReq) (resp *types.CollectionResp, err error) {
	// 从token中获取会员id
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	member, err := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员信息失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 查询是否已经收藏
	collectionList, err := l.svcCtx.MemberProductCollectionService.QueryMemberProductCollectionList(l.ctx, &umsclient.QueryMemberProductCollectionListReq{
		PageNum:   1,
		PageSize:  10,
		MemberId:  memberId,
		ProductId: req.ProductId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员是否已经收藏失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 如果查询不到收藏记录,则添加
	if len(collectionList.List) == 0 {
		_, err = l.svcCtx.MemberProductCollectionService.AddMemberProductCollection(l.ctx, &umsclient.AddMemberProductCollectionReq{
			MemberId:        member.MemberId,
			MemberNickName:  member.Nickname,
			MemberIcon:      member.Avatar,
			ProductId:       req.ProductId,
			ProductName:     req.ProductName,
			ProductPic:      req.ProductPic,
			ProductSubTitle: req.ProductSubTitle,
			ProductPrice:    req.ProductPrice,
		})

		if err != nil {
			logc.Errorf(l.ctx, "添加商品收藏失败,参数: %+v,异常：%s", req, err.Error())
			s, _ := status.FromError(err)
			return nil, errorx.NewDefaultError(s.Message())
		}
	}

	return &types.CollectionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
