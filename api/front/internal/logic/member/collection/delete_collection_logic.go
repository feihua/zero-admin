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

// DeleteCollectionLogic 删除商品收藏
/*
Author: LiuFeiHua
Date: 2025/6/19 10:58
*/
type DeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCollection 删除商品收藏
func (l *DeleteCollectionLogic) DeleteCollection(req *types.CollectionDeleteReq) (resp *types.CollectionResp, err error) {
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

	return &types.CollectionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
