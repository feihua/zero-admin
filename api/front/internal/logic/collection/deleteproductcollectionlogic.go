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

// DeleteProductCollectionLogic 删除商品收藏
/*
Author: LiuFeiHua
Date: 2024/5/16 13:46
*/
type DeleteProductCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCollectionLogic {
	return &DeleteProductCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductCollection 删除商品收藏
func (l *DeleteProductCollectionLogic) DeleteProductCollection(req *types.ProductCollectionDeleteReq) (resp *types.ProductCollectionDeleteResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberProductCollectionService.DeleteMemberProductCollection(l.ctx, &umsclient.DeleteMemberProductCollectionReq{
		Ids:      req.Ids,
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
