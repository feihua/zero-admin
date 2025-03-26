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

// ClearProductCollectionLogic 清空当前用户商品收藏列表
/*
Author: LiuFeiHua
Date: 2025/3/26 17:35
*/
type ClearProductCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearProductCollectionLogic {
	return &ClearProductCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ClearProductCollection 清空当前用户商品收藏列表
func (l *ClearProductCollectionLogic) ClearProductCollection() (resp *types.ProductCollectionDeleteResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberProductCollectionService.DeleteMemberProductCollection(l.ctx, &umsclient.DeleteMemberProductCollectionReq{
		Ids:      nil,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品收藏失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ProductCollectionDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
