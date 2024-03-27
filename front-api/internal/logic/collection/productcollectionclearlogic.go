package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCollectionClearLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:31
*/
type ProductCollectionClearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCollectionClearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCollectionClearLogic {
	return &ProductCollectionClearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCollectionClear 清空收藏记录
func (l *ProductCollectionClearLogic) ProductCollectionClear() (resp *types.ProductCollectionClearResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberProductCollectionService.MemberProductCollectionDelete(l.ctx, &umsclient.MemberProductCollectionDeleteReq{
		MemberId: memberId,
	})

	return &types.ProductCollectionClearResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
