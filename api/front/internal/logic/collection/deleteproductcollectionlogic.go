package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductCollectionLogic 删除商品收藏/清空当前用户商品收藏列表
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

// DeleteProductCollection 删除商品收藏/清空当前用户商品收藏列表
func (l *DeleteProductCollectionLogic) DeleteProductCollection(req *types.ProductCollectionDeleteReq) (resp *types.ProductCollectionDeleteResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberProductCollectionService.MemberProductCollectionDelete(l.ctx, &umsclient.MemberProductCollectionDeleteReq{
		Ids:      req.Ids,
		MemberId: memberId,
	})

	return &types.ProductCollectionDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
